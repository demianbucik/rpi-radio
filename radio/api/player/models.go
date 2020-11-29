package player

import "radio/api/tracks"

type StateDto struct {
	Tracks          []*tracks.TrackDto `json:"tracks,omitempty"`
	NTracks         *int               `json:"nTracks,omitempty"`
	CurrentTrack    *int               `json:"currentTrack,omitempty"`
	CurrentPosition *float32           `json:"currentPosition,omitempty"`
	Volume          *int               `json:"volume,omitempty"`
	MediaState      *string            `json:"mediaState,omitempty"`
}
