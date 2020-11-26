package models

import (
	"time"
)

type Track struct {
	ID             uint
	Name           string
	Url            string
	Thumbnail      string
	PlaylistTracks []*PlaylistTrack
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Playlist struct {
	ID             uint
	Name           string
	PlaylistTracks []*PlaylistTrack
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type PlaylistTrack struct {
	ID         uint
	PlaylistID uint
	Playlist   *Playlist
	TrackID    uint
	Track      *Track
	Position   *uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
