package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"

	"radio/api/playlists"
	"radio/api/tracks"
	"radio/api/utils"
)

func NewRouter(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(utils.RequestCtx)
	router.Use(utils.Logger)
	router.Use(utils.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	playlistsApi := playlists.New(db)
	tracksApi := tracks.New(db)

	playlistRoutes(router, playlistsApi)
	tracksRoutes(router, tracksApi)

	return router
}

func playlistRoutes(router *chi.Mux, playlistsApi *playlists.Api) {
	router.Route("/playlists", func(r chi.Router) {
		r.Get("/", playlistsApi.List)
		r.Post("/", playlistsApi.Create)

		r.Route("/{playlistId:[0-9]+}", func(r chi.Router) {
			r.Use(playlistsApi.PlaylistCtx)
			r.Get("/", playlistsApi.Get)
			r.Put("/", playlistsApi.Update)
			r.Delete("/", playlistsApi.Delete)

			r.Route("/tracks", func(r chi.Router) {
				r.Post("/", playlistsApi.AddTracks)
				r.Put("/", playlistsApi.ReorderTracks)

				r.Route("/{position:[0-9]+}", func(r chi.Router) {
					r.Use(playlistsApi.PlaylistTrackCtx)
					r.Delete("/", playlistsApi.DeleteTrack)
				})
			})
		})
	})
}

func tracksRoutes(router *chi.Mux, tracksApi *tracks.Api) {
	router.Route("/tracks", func(r chi.Router) {
		r.Get("/", tracksApi.List)
		r.Post("/", tracksApi.Create)

		r.Route("/{trackId:[0-9]+}", func(r chi.Router) {
			r.Use(tracksApi.TrackCtx)
			r.Put("/", tracksApi.Update)
			r.Delete("/", tracksApi.Delete)
		})
	})
}
