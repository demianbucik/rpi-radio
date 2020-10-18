package tracks

import "radio/models"

func ToDto(track *models.Track) *TrackDto {
	return &TrackDto{
		ID:        track.ID,
		Name:      track.Name,
		Url:       track.Url,
		Thumbnail: track.Thumbnail,
	}
}

func FromDto(dto *TrackDto) *models.Track {
	track := &models.Track{
		Name:      dto.Name,
		Url:       dto.Url,
		Thumbnail: dto.Thumbnail,
	}
	track.ID = dto.ID
	return track
}

func ToDtos(tracks []*models.Track) []*TrackDto {
	var dtos []*TrackDto
	for _, track := range tracks {
		dtos = append(dtos, ToDto(track))
	}
	return dtos
}

func FromDtos(dtos []*TrackDto) []*models.Track {
	var tracks []*models.Track
	for _, dto := range dtos {
		tracks = append(tracks, FromDto(dto))
	}
	return tracks
}

func ToTrackDto(plTrack *models.PlaylistTrack) *TrackDto {
	dto := ToDto(plTrack.Track)
	dto.Position = plTrack.Position
	return dto
}

func ToTrackDtos(plTracks []*models.PlaylistTrack) []*TrackDto {
	var dtos []*TrackDto
	for _, plTrack := range plTracks {
		dtos = append(dtos, ToTrackDto(plTrack))
	}
	return dtos
}
