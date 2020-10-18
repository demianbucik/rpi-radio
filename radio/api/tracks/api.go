package tracks

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
	var tracks []*models.Track

	if err := a.db.Find(&tracks).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, ToDtos(tracks))
}

func (a *Api) Create(w http.ResponseWriter, r *http.Request) {
	var dto TrackDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	track := FromDto(&dto)

	if err := a.db.Create(track).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusCreated, ToDto(track))
}

func (a *Api) Update(w http.ResponseWriter, r *http.Request) {
	var dto TrackDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.ServerError(w, r, err)
		return
	}

	trackId := r.Context().Value("track").(*models.Track).ID

	track := FromDto(&dto)
	track.ID = trackId

	if err := a.db.Model(track).Updates(track).Error; err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, ToDto(track))
}

var errTrackProtected = errors.New("cannot delete, track is used in a playlist")

func (a *Api) Delete(w http.ResponseWriter, r *http.Request) {
	track := r.Context().Value("track").(*models.Track)

	err := a.db.Transaction(func(tx *gorm.DB) error {
		var pt []*models.PlaylistTrack
		if err := tx.Find(&pt, "track_id = ?", track.ID).Error; err != nil {
			return err
		}
		if len(pt) > 0 {
			return errTrackProtected
		}

		if err := tx.Delete(&models.Track{}, track.ID).Error; err != nil {
			return err
		}
		return nil
	})
	if err == errTrackProtected {
		utils.BadRequest(w, r, err)
		return
	}
	if err != nil {
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, nil)
}
