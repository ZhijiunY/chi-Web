package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/ZhijiunY/chi-web/internal/config"
	handlers "github.com/ZhijiunY/chi-web/internal/handlers"
	"github.com/ZhijiunY/chi-web/models"

	"github.com/alexedwards/scs/v2"
)

var SessionManager *scs.SessionManager

var app config.AppConfig

func main() {

	// store article in session
	gob.Register(models.Article{})

	SessionManager = scs.New()
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.Secure = false
	SessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = SessionManager

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		// log.Fatalf("Error starting server: %s", err)
		log.Fatal(err)

	}
}
