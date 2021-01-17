package player

import (
	"errors"
	"sync"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/apex/log"

	"radio/common/youtube"
	"radio/models"
)

var (
	ErrNoTracks                = errors.New("no tracks to play")
	ErrEndReached              = errors.New("reached the end")
	ErrAlreadyPlaying          = errors.New("player is already playing")
	ErrAlreadyPaused           = errors.New("player is already paused")
	ErrMediaNotPlayingOrPaused = errors.New("media not playing or paused")
	ErrInvalidPosition         = errors.New("invalid track position")
)

var events = []vlc.Event{
	vlc.MediaPlayerTimeChanged,
	vlc.MediaPlayerEndReached,
}

type Player struct {
	mu sync.RWMutex

	player       *vlc.Player
	eventManager *vlc.EventManager

	youtube *youtube.Client

	vlcEventsChan chan *vlcEvent

	trackIdx int
	tracks   []*models.Track

	release release
}

type State struct {
	Tracks          []*models.Track
	CurrentPosition *int
	CurrentTime     *float32
	Volume          *int
	MediaState      *vlc.MediaState
}

type release struct {
	media    []*vlc.Media
	eventIDs []vlc.EventID
}

func New() (*Player, error) {
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		return nil, err
	}
	vlcPlayer, err := vlc.NewPlayer()
	if err != nil {
		return nil, err
	}
	manager, err := vlcPlayer.EventManager()
	if err != nil {
		return nil, err
	}

	player := &Player{
		player:        vlcPlayer,
		eventManager:  manager,
		youtube:       youtube.NewClient(),
		vlcEventsChan: make(chan *vlcEvent, 1),
	}

	if err := player.registerEvents(); err != nil {
		return nil, err
	}

	go player.loop()

	return player, nil
}

func (p *Player) registerEvents() error {
	eventCallback := func(event vlc.Event, userData interface{}) {
		p.vlcEventsChan <- &vlcEvent{event: event, userData: userData}
	}

	for _, event := range events {
		eventID, err := p.eventManager.Attach(event, eventCallback, nil)
		if err != nil {
			return err
		}
		p.release.eventIDs = append(p.release.eventIDs, eventID)
	}

	return nil
}

func (p *Player) GetState(includeTracks bool) *State {
	p.mu.Lock()
	defer p.mu.Unlock()

	state := &State{
		CurrentPosition: &p.trackIdx,
	}

	if mediaState, err := p.player.MediaState(); err == nil {
		state.MediaState = &mediaState
	} else {
		log.WithError(err).Warn("GetState: failed to get media state")
	}

	if time, err := p.player.MediaPosition(); err == nil {
		state.CurrentTime = &time
	} else {
		log.WithError(err).Warn("GetState: failed to get media time")
	}

	if vol, err := p.player.Volume(); err == nil {
		state.Volume = &vol
	} else {
		log.WithError(err).Warn("GetState: failed to get volume")
	}

	if includeTracks {
		state.Tracks = make([]*models.Track, len(p.tracks))
		for i, track := range p.tracks {
			t := *track
			state.Tracks[i] = &t
		}
	}

	return state
}

func (p *Player) SetVolume(volume int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.player.SetVolume(volume)
}

func (p *Player) Play() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.player.IsPlaying() {
		return ErrAlreadyPlaying
	}

	media, _ := p.player.Media()
	if media == nil {
		return p.loadTrackAndPlay()
	}

	return p.player.Play()
}

func (p *Player) Pause() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.player.IsPlaying() {
		return ErrAlreadyPaused
	}

	return p.player.SetPause(true)
}

func (p *Player) Stop() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.player.Stop()
}

func (p *Player) PlayNext() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.playNext()
}

func (p *Player) PlayPrevious() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.playPrevious()
}

func (p *Player) SetTime(percentage float32) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	mediaState, err := p.player.MediaState()
	if err != nil {
		return err
	}
	if !(mediaState == vlc.MediaPlaying || mediaState == vlc.MediaPaused) {
		return ErrMediaNotPlayingOrPaused
	}

	return p.player.SetMediaPosition(percentage)
}

func (p *Player) EnqueueTracks(override bool, tracks ...*models.Track) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if override {
		p.tracks = tracks
		p.trackIdx = 0
		return p.loadTrackAndPlay()
	}

	p.tracks = append(p.tracks, tracks...)

	return nil
}

func (p *Player) RemoveAllTracks() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.player.Stop(); err != nil {
		log.WithError(err).Warn("RemoveAllTracks: failed to stop the player")
	}

	p.trackIdx = 0
	for i := range p.tracks {
		p.tracks[i] = nil
	}
	p.tracks = p.tracks[:0]
}

func (p *Player) DeleteTrack(position int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if position < 0 || position >= len(p.tracks) {
		return ErrInvalidPosition
	}

	n := len(p.tracks)
	copy(p.tracks[position:], p.tracks[position+1:])
	p.tracks[n-1] = nil
	p.tracks = p.tracks[:n-1]

	if p.trackIdx == position {
		isPlaying := p.player.IsPlaying()
		if err := p.player.Stop(); err != nil {
			log.WithError(err).Warn("DeleteTrack: failed to stop the player")
		}
		if isPlaying && len(p.tracks) > 0 {
			return p.loadTrackAndPlay()
		}
		return nil
	}

	if p.trackIdx > position {
		p.trackIdx--
	}

	return nil
}

func (p *Player) loadTrackAndPlay() error {
	if len(p.tracks) == 0 {
		return ErrNoTracks
	}
	track := p.tracks[p.trackIdx]
	url, err := p.youtube.GetBestAudioStreamURL(track.Url)
	if err != nil {
		return err
	}

	// If any, previous media will be released
	// https://www.videolan.org/developers/vlc/doc/doxygen/html/group__libvlc__media__player.html#gadeb7ac440f41dbb2aa1a7811904099b1
	_, err = p.player.LoadMediaFromURL(url)
	if err != nil {
		return err
	}

	return p.player.Play()
}

func (p *Player) playNext() error {
	if len(p.tracks) == 0 {
		return ErrNoTracks
	}
	if p.trackIdx == len(p.tracks)-1 {
		return ErrEndReached
	}

	p.trackIdx++
	return p.loadTrackAndPlay()
}

func (p *Player) playPrevious() error {
	if len(p.tracks) == 0 {
		return ErrNoTracks
	}
	if p.trackIdx == 0 {
		return ErrEndReached
	}

	p.trackIdx--
	return p.loadTrackAndPlay()
}

func (p *Player) loop() {
	for {
		event := <-p.vlcEventsChan
		switch event.event {
		case vlc.MediaPlayerEndReached:
			if err := p.PlayNext(); err != nil {
				if err == ErrEndReached {
					if err := p.Stop(); err != nil {
						log.WithError(err).Debug("loop: could not stop when end was reached")
					}
					continue
				}
				log.WithError(err).Error("loop: error while trying to play the next track")
			}
		default:
		}
	}
}

type vlcEvent struct {
	event    vlc.Event
	userData interface{}
}

func (p *Player) Shutdown() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, eventID := range p.release.eventIDs {
		p.eventManager.Detach(eventID)
	}
	if media, err := p.player.Media(); err == nil {
		if err := media.Release(); err != nil {
			log.WithError(err).Debug("Shutdown: media release failed")
		}
	}

	if err := p.player.Stop(); err != nil {
		log.WithError(err).Debug("Shutdown: player stop failed")
	}
	if err := p.player.Release(); err != nil {
		log.WithError(err).Debug("Shutdown: player release failed")
	}
	if err := vlc.Release(); err != nil {
		log.WithError(err).Debug("Shutdown: VLC release failed")
	}
}
