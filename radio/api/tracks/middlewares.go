package tracks

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"radio/api/utils"
	"radio/models"
)

func (a *Api) TrackCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "trackId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		var track models.Track
		if err := a.db.First(&track, id).Error; err != nil {
			utils.BadRequest(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), utils.TrackCtxKey, &track)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
