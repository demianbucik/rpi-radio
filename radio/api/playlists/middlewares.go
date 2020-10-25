package playlists

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"radio/api/utils"
	"radio/models"
)

func (a *Api) PlaylistCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "playlistId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		var playlist models.Playlist
		if err := a.db.First(&playlist, id).Error; err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "playlist", &playlist)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Api) PlaylistTrackCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		playlist := r.Context().Value("playlist").(*models.Playlist)

		positionStr := chi.URLParam(r, "position")
		position, err := strconv.Atoi(positionStr)
		if err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		var playlistTrack models.PlaylistTrack
		err = a.db.First(&playlistTrack, "playlist_id = ? AND position = ?", playlist.ID, position).Error
		if err != nil {
			utils.ServerError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "playlistTrack", &playlistTrack)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
