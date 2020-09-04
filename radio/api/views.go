package api

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"radio/api/playlists"
	"radio/api/tracks"
)

func Init(db *gorm.DB, router *chi.Mux) {

	playlistsApi := playlists.New(db)
	tracksApi := tracks.New(db)

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
