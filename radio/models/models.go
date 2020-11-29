package models

import (
	"time"
)

type Track struct {
	ID             int
	Name           string
	Url            string
	Thumbnail      string
	PlaylistTracks []*PlaylistTrack
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Playlist struct {
	ID             int
	Name           string
	PlaylistTracks []*PlaylistTrack
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (p *Playlist) GetTracks() []*Track {
	var ts []*Track
	for _, pt := range p.PlaylistTracks {
		ts = append(ts, pt.Track)
	}
	return ts
}

type PlaylistTrack struct {
	ID         int
	PlaylistID int
	Playlist   *Playlist
	TrackID    int
	Track      *Track
	Position   *int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
