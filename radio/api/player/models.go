package player

import "radio/api/tracks"

type StateDto struct {
	Tracks          []*tracks.TrackDto `json:"tracks,omitempty"`
	NTracks         *int               `json:"nTracks,omitempty"`
	CurrentPosition *int               `json:"currentPosition,omitempty"`
	CurrentTime     *float32           `json:"currentTime,omitempty"`
	Volume          *int               `json:"volume,omitempty"`
	MediaState      *string            `json:"mediaState,omitempty"`
}
