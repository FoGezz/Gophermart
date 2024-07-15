package main

import (
	"Gophermart/cmd/gophermart/config"
	"Gophermart/internal/app/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	cfg := new(config.Config)
	cfg.Load()
	log.Default().Printf("Running on conf: %s", cfg)
	db, err := sqlx.Connect("postgres", cfg.DBDSN)
	if err != nil {
		log.Fatalln(err)
	}
	app := config.NewApp(cfg, db)

	mux := chi.NewRouter()
	handler.InitRoutes(mux, app)
	listenErr := http.ListenAndServe(":8080", mux)
	if listenErr != nil {
		return
	}
}
