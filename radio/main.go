package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"radio/api"
	"radio/models"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	db, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	api.Init(db, router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("ListenAndServe failed", err)
		}
	}()

	log.Println("ListenAndServe")

	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	<-shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown failed", err)
	}
	log.Println("Server shutdown")
}
