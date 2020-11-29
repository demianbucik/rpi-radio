package playlists

import (
	"errors"
	"net/http"

	"github.com/segmentio/encoding/json"
	"gorm.io/gorm"

	"radio/api/utils"
	"radio/models"
)

type Api struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Api {
	return &Api{db: db}
}

func (a *Api) List(w http.ResponseWriter, r *http.Request) {
	var playlists []*models.Playlist

	err := a.db.Preload("PlaylistTracks", func(db *gorm.DB) *gorm.DB { return db.Order("position") }).
		Preload("PlaylistTracks.Track").Find(&playlists).Error

	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, PlaylistsToDtos(playlists))
}

func (a *Api) Get(w http.ResponseWriter, r *http.Request) {
	var playlist models.Playlist

	id := r.Context().Value(utils.PlaylistCtxKey).(*models.Playlist).ID

	err := a.db.Preload("PlaylistTracks", func(db *gorm.DB) *gorm.DB { return db.Order("position") }).
		Preload("PlaylistTracks.Track").Find(&playlist, id).Error

	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, PlaylistToDto(&playlist))
}

func (a *Api) Create(w http.ResponseWriter, r *http.Request) {
	var dto PlaylistDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	playlist := DtoToPlaylist(&dto)

	if err := a.db.Create(playlist).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusCreated, PlaylistToDto(playlist))
}

func (a *Api) Update(w http.ResponseWriter, r *http.Request) {
	var dto PlaylistDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	id := r.Context().Value(utils.PlaylistCtxKey).(*models.Playlist).ID

	playlist := DtoToPlaylist(&dto)
	playlist.ID = id

	if err := a.db.Model(playlist).Updates(playlist).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, PlaylistToDto(playlist))
}

func (a *Api) Delete(w http.ResponseWriter, r *http.Request) {
	playlist := r.Context().Value(utils.PlaylistCtxKey).(*models.Playlist)

	err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.PlaylistTrack{}, "playlist_id = ?", playlist.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&playlist).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) AddTracks(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Tracks   []int `json:"tracks"`
		Position *int  `json:"position"`
	}

	// TODO: refactor

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	playlist := r.Context().Value(utils.PlaylistCtxKey).(*models.Playlist)

	err := a.db.Transaction(func(tx *gorm.DB) error {
		nTracks := len(dto.Tracks)

		var count64 int64
		err := tx.Model(&models.PlaylistTrack{}).Where("playlist_id = ?", playlist.ID).Count(&count64).Error
		if err != nil {
			return err
		}

		count := int(count64)
		pos := count

		if dto.Position != nil {
			pos = *dto.Position
			if pos > count {
				return errors.New("invalid position")
			}

			err := tx.Model(&models.PlaylistTrack{}).Where("position >= ?", pos).
				UpdateColumn("position", gorm.Expr("position + ?", nTracks)).Error
			if err != nil {
				return err
			}
		}

		var pts []*models.PlaylistTrack

		for index, id := range dto.Tracks {
			newPos := pos + index
			pts = append(pts, &models.PlaylistTrack{
				PlaylistID: playlist.ID,
				TrackID:    id,
				Position:   &newPos,
			})
		}

		if err := tx.Create(&pts).Error; err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusCreated, nil)
}

// Reorder tracks
func (a *Api) ReorderTracks(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Position     int `json:"position"`
		InsertBefore int `json:"insertBefore"`
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	playlist := r.Context().Value(utils.PlaylistCtxKey).(*models.Playlist)

	var pt models.PlaylistTrack

	if err := a.db.Find(&pt, "playlist_id = ? AND position = ?", playlist.ID, dto.Position).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	err := a.db.Transaction(func(tx *gorm.DB) error {
		*pt.Position = dto.InsertBefore
		var err error

		if dto.Position < dto.InsertBefore {
			*pt.Position--
			err = tx.Model(&models.PlaylistTrack{}).Where("position > ? AND position < ?", dto.Position, dto.InsertBefore).
				UpdateColumn("position", gorm.Expr("position - 1")).Error
		} else {
			err = tx.Model(&models.PlaylistTrack{}).Where("position >= ? AND position < ?", dto.InsertBefore, dto.Position).
				UpdateColumn("position", gorm.Expr("position + 1")).Error
		}

		if err != nil {
			return err
		}

		return tx.Model(pt).Updates(pt).Error
	})

	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) DeleteTrack(w http.ResponseWriter, r *http.Request) {
	pt := r.Context().Value(utils.PlaylistTrackCtxKey).(*models.PlaylistTrack)

	err := a.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&pt).Error
		if err != nil {
			return err
		}
		err = tx.Model(&models.PlaylistTrack{}).Where("position > ?", pt.Position).
			UpdateColumn("position", gorm.Expr("position - 1")).Error
		return err
	})

	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, nil)
}
