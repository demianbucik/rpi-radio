package player

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"radio/api/tracks"
	"radio/api/utils"
	"radio/common/player"
)

type Api struct {
	db     *gorm.DB
	player *player.Player
}

func New(db *gorm.DB, p *player.Player) *Api {
	return &Api{
		db:     db,
		player: p,
	}
}

func (a *Api) State(w http.ResponseWriter, r *http.Request) {
	tracksStr := r.URL.Query().Get("tracks")
	includeTracks, _ := strconv.ParseBool(tracksStr)

	state := a.player.GetState(includeTracks)
	utils.Respond(w, http.StatusOK, StateToDto(state))
}

func (a *Api) DeleteTrack(w http.ResponseWriter, r *http.Request) {
	positionStr := chi.URLParam(r, "position")
	position, _ := strconv.Atoi(positionStr)

	if err := a.player.DeleteTrack(position - 1); err != nil {
		if err == player.ErrInvalidPosition {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) Play(w http.ResponseWriter, r *http.Request) {
	if err := a.player.Play(); err != nil {
		if err == player.ErrAlreadyPlaying {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) Pause(w http.ResponseWriter, r *http.Request) {
	if err := a.player.Pause(); err != nil {
		if err == player.ErrAlreadyPaused {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) Stop(w http.ResponseWriter, r *http.Request) {
	if err := a.player.Stop(); err != nil {
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) Next(w http.ResponseWriter, r *http.Request) {
	if err := a.player.PlayNext(); err != nil {
		if err == player.ErrEndReached || err == player.ErrNoTracks {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) Previous(w http.ResponseWriter, r *http.Request) {
	if err := a.player.PlayPrevious(); err != nil {
		if err == player.ErrEndReached || err == player.ErrNoTracks {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) SetTime(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		TimePercent float32 `json:"timePercent"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.BadRequest(w, r, err)
		return
	}
	if err := a.player.SetTime(dto.TimePercent); err != nil {
		if err == player.ErrMediaNotPlayingOrPaused {
			utils.BadRequest(w, r, err)
			return
		}
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) SetVolume(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Volume int `json:"volume"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utils.BadRequest(w, r, err)
		return
	}
	if err := a.player.SetVolume(dto.Volume); err != nil {
		utils.ServerError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (a *Api) enqueueTracksFunc(override bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto []*tracks.TrackDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		ts := tracks.DtosToTracks(dto)
		if err := a.player.EnqueueTracks(override, ts...); err != nil {
			utils.ServerError(w, r, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}

func (a *Api) EnqueueTracks(w http.ResponseWriter, r *http.Request) {
	a.enqueueTracksFunc(false)(w, r)
}

func (a *Api) EnqueueTracksOverride(w http.ResponseWriter, r *http.Request) {
	a.enqueueTracksFunc(true)(w, r)
}
