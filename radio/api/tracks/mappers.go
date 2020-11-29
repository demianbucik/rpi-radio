package tracks

import (
	"radio/models"
)

func TrackToDto(track *models.Track) *TrackDto {
	return &TrackDto{
		ID:        track.ID,
		Name:      track.Name,
		Url:       track.Url,
		Thumbnail: track.Thumbnail,
	}
}

func DtoToTrack(dto *TrackDto) *models.Track {
	return &models.Track{
		ID:        dto.ID,
		Name:      dto.Name,
		Url:       dto.Url,
		Thumbnail: dto.Thumbnail,
	}
}

func TracksToDtos(tracks []*models.Track) []*TrackDto {
	var dtos []*TrackDto
	for _, track := range tracks {
		dtos = append(dtos, TrackToDto(track))
	}
	return dtos
}

func DtosToTracks(dtos []*TrackDto) []*models.Track {
	var tracks []*models.Track
	for _, dto := range dtos {
		tracks = append(tracks, DtoToTrack(dto))
	}
	return tracks
}

func PlaylistTrackToDto(plTrack *models.PlaylistTrack) *TrackDto {
	dto := TrackToDto(plTrack.Track)
	dto.Position = plTrack.Position
	return dto
}

func PlaylistTracksToDtos(plTracks []*models.PlaylistTrack) []*TrackDto {
	var dtos []*TrackDto
	for _, t := range plTracks {
		dtos = append(dtos, PlaylistTrackToDto(t))
	}
	return dtos
}
