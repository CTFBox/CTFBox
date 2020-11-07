package main

import (
	"context"
	"github.com/CTFBox/CTFBox/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CTFBox/CTFBox/router"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

var (
	SESSION_KEY = []byte(os.Getenv("SESSION_KEY")[:])
)

func main() {
	db, err := SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger, _ := zap.NewDevelopment()
	repo, err := repository.NewGormRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	handler := &router.Handlers{
		Repo:       repo,
		Logger:     logger,
		SessionKey: SESSION_KEY,
		SessionOption: sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
		Origin: os.Getenv("ORIGIN"),
	}

	e := handler.SetupRoute(db)

	// サーバースタート
	go func() {
		if err := e.Start(":3000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
