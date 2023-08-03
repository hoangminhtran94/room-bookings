package main

import (
	"net/http"

	"github.com/hoangminhtran94/room-bookings/pkg/config"
	"github.com/hoangminhtran94/room-bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(WriteToConsole)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/make-reservation", handlers.Repo.Resevation)
	router.Get("/check-availability", handlers.Repo.Availability)
	router.Post("/check-availability", handlers.Repo.PostAvailability)
	router.Get("/generals-quarters", handlers.Repo.Generals)
	router.Get("/majors-suite", handlers.Repo.MajorSuite)
	router.Get("/contacts",handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
