package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hoangminhtran94/room-bookings/internal/config"
	"github.com/hoangminhtran94/room-bookings/internal/handlers"
	"github.com/hoangminhtran94/room-bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":3080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change to true in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	template, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = template
	app.UseCache = false
	//Passed app config to handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("Starting server at" + portNumber)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
