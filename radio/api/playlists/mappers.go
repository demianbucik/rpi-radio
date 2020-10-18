package playlists

import (
	"radio/api/tracks"
	"radio/models"
)

func ToDto(playlist *models.Playlist) *PlaylistDto {
	dto := &PlaylistDto{
		ID:     playlist.ID,
		Name:   playlist.Name,
		Tracks: tracks.ToTrackDtos(playlist.PlaylistTracks),
	}
	return dto
}

func FromDto(dto *PlaylistDto) *models.Playlist {
	playlist := &models.Playlist{
		Name: dto.Name,
		// Tracks: tracks.FromDtos(dto.Tracks),
	}
	playlist.ID = dto.ID
	return playlist
}

func ToDtos(playlists []*models.Playlist) []*PlaylistDto {
	var dtos []*PlaylistDto
	for _, playlist := range playlists {
		dtos = append(dtos, ToDto(playlist))
	}
	return dtos
}

func FromDtos(dtos []*PlaylistDto) []*models.Playlist {
	var playlists []*models.Playlist
	for _, dto := range dtos {
		playlists = append(playlists, FromDto(dto))
	}
	return playlists
}

func GetTracks(playlist *models.Playlist) []*models.Track {
	var ts []*models.Track
	for _, pt := range playlist.PlaylistTracks {
		ts = append(ts, pt.Track)
	}
	return ts
}
