package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"banachtech.github.com/weirdville/cmd/pkg/config"
	"banachtech.github.com/weirdville/cmd/pkg/handlers"
	"banachtech.github.com/weirdville/cmd/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	// create session
	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // No need for https for testing mode

	// set config session
	app.Session = session

	// create template cache
	ts, err := render.CreateTemplate()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = ts

	// initialise templates with cache
	render.NewTemplate(&app)

	// initialise repo for handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// instantiate server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// run server
	fmt.Println("Starting application on port " + portNumber)
	srv.ListenAndServe()
}
