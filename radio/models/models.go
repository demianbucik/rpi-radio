package models

import (
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Name           string
	Url            string
	Thumbnail      string
	PlaylistTracks []PlaylistTrack
}

type Playlist struct {
	gorm.Model
	Name           string
	PlaylistTracks []PlaylistTrack
}

type PlaylistTrack struct {
	gorm.Model
	PlaylistID uint
	Playlist   *Playlist
	TrackID    uint
	Track      *Track
	Position   *uint
}
