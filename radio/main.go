package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/apex/log"

	"radio/api"
	"radio/app"
	"radio/app/config"
	"radio/common/player"
	"radio/models"
)

func main() {
	app.Startup()

	db, err := models.NewDB()
	if err != nil {
		log.WithError(err).Fatal("Opening DB failed")
	}

	p, err := player.New()
	if err != nil {
		log.WithError(err).Fatal("Opening VLC player failed")
	}

	router := api.NewRouter(db, p)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Env.PORT),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.WithError(err).Fatal("ListenAndServe failed")
		}
	}()

	waitForShutdown(srv, p)
}

func waitForShutdown(srv *http.Server, p *player.Player) {

	log.Info("Listen and serve")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		log.WithError(err).Fatal("Server shutdown failed")
	}
	p.Shutdown()
	log.Info("Server shutdown")
}
