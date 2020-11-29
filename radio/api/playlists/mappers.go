package playlists

import (
	"radio/api/tracks"
	"radio/models"
)

func PlaylistToDto(playlist *models.Playlist) *PlaylistDto {
	dto := &PlaylistDto{
		ID:     playlist.ID,
		Name:   playlist.Name,
		Tracks: tracks.PlaylistTracksToDtos(playlist.PlaylistTracks),
	}
	return dto
}

func DtoToPlaylist(dto *PlaylistDto) *models.Playlist {
	return &models.Playlist{
		ID:   dto.ID,
		Name: dto.Name,
		// Tracks: tracks.DtosToTracks(dto.Tracks),
	}
}

func PlaylistsToDtos(playlists []*models.Playlist) []*PlaylistDto {
	var dtos []*PlaylistDto
	for _, playlist := range playlists {
		dtos = append(dtos, PlaylistToDto(playlist))
	}
	return dtos
}

func DtosToPlaylists(dtos []*PlaylistDto) []*models.Playlist {
	var playlists []*models.Playlist
	for _, dto := range dtos {
		playlists = append(playlists, DtoToPlaylist(dto))
	}
	return playlists
}
