package api

import (
	"compress/gzip"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"

	"radio/api/player"
	"radio/api/playlists"
	"radio/api/tracks"
	"radio/api/utils"
	"radio/app/config"
	commonPlayer "radio/common/player"
)

func NewRouter(db *gorm.DB, p *commonPlayer.Player) *chi.Mux {
	router := chi.NewRouter()
	router.Use(utils.RequestCtx)
	router.Use(utils.CORS)
	router.Use(utils.Logger)
	router.Use(utils.Recoverer)
	router.Use(middleware.Compress(gzip.DefaultCompression))
	router.Use(middleware.Heartbeat("/ping"))

	playlistsApi := playlists.New(db)
	tracksApi := tracks.New(db)
	playerApi := player.New(db, p)

	router.Route("/api", func(r chi.Router) {
		playlistRoutes(r, playlistsApi)
		tracksRoutes(r, tracksApi)
		playerRoutes(r, playerApi)
	})

	fs := http.FileServer(http.Dir(config.Env.STATIC_FILES_DIR))
	router.Handle("/*", fs)

	return router
}

func playerRoutes(router chi.Router, playerApi *player.Api) {
	router.Route("/player", func(r chi.Router) {
		r.Post("/play", playerApi.Play)
		r.Post("/pause", playerApi.Pause)
		r.Post("/stop", playerApi.Stop)
		r.Post("/next", playerApi.Next)
		r.Post("/previous", playerApi.Previous)

		r.Post("/tracks", playerApi.EnqueueTracksOverride)
		r.Put("/tracks", playerApi.EnqueueTracks)
		r.Delete("/tracks/{position:[0-9]+}", playerApi.DeleteTrack)

		r.Put("/position", playerApi.SetPosition)
		r.Put("/volume", playerApi.SetVolume)

		r.Get("/state", playerApi.State)
	})
}

func playlistRoutes(router chi.Router, playlistsApi *playlists.Api) {
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

func tracksRoutes(router chi.Router, tracksApi *tracks.Api) {
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
