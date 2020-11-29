package playlists

import "radio/api/tracks"

type PlaylistDto struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Tracks []*tracks.TrackDto `json:"tracks,omitempty"`
}
