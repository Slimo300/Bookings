package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Slimo300/Bookings/internal/config"
	"github.com/Slimo300/Bookings/internal/handlers"
	"github.com/Slimo300/Bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig
var session scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error when initializing config: %v", err)
	}

	server := http.Server{
		Handler: routes(),
		Addr:    PORT,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func run() error {
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = *scs.New()
	session.Lifetime = 24 * time.Second
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = &session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo()
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	return nil
}
