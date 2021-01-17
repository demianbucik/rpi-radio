package player

import (
	"radio/api/tracks"
	"radio/common/player"
)

func StateToDto(state *player.State) *StateDto {
	dto := &StateDto{
		Tracks:          tracks.TracksToDtos(state.Tracks),
		CurrentPosition: state.CurrentPosition,
		CurrentTime:     state.CurrentTime,
		Volume:          state.Volume,
	}
	for i, track := range dto.Tracks {
		pos := i + 1
		track.Position = &pos
	}
	if state.Tracks != nil {
		nTracks := len(state.Tracks)
		dto.NTracks = &nTracks
	}
	if state.MediaState != nil {
		mediaState := mediaStateDescription(*state.MediaState)
		dto.MediaState = &mediaState
	}
	return dto
}
