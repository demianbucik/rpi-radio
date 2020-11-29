package player

import vlc "github.com/adrg/libvlc-go/v3"

const (
	MediaStateNothingSpecial string = "nothingSpecial"
	MediaStateOpening        string = "opening"
	MediaStateBuffering      string = "buffering"
	MediaStatePlaying        string = "playing"
	MediaStatePaused         string = "paused"
	MediaStateStopped        string = "stopped"
	MediaStateEnded          string = "ended"
	MediaStateError          string = "error"
)

func mediaStateDescription(state vlc.MediaState) string {
	switch state {
	case vlc.MediaOpening:
		return MediaStateOpening
	case vlc.MediaBuffering:
		return MediaStateBuffering
	case vlc.MediaPlaying:
		return MediaStatePlaying
	case vlc.MediaPaused:
		return MediaStatePaused
	case vlc.MediaStopped:
		return MediaStateStopped
	case vlc.MediaEnded:
		return MediaStateEnded
	case vlc.MediaError:
		return MediaStateError
	default:
		return MediaStateNothingSpecial
	}
}
