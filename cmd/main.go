package main

import (
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/privetyaonton/crud-api-go/config"
	"github.com/privetyaonton/crud-api-go/internal/repo/postgres"
	"github.com/privetyaonton/crud-api-go/internal/restapi"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config new failed: %v", err)
	}
	pg, err := postgres.New(cfg.Postgres.GetConnectionUrl())
	if err != nil {
		log.Fatalf("postgres new failed: %v", err)
	}

	app := restapi.New(pg)

	defer app.Close()

	mux := http.NewServeMux()
	mux.Handle("/users", http.HandlerFunc(app.CheckMethod))

	log.Fatal(http.ListenAndServe(cfg.Addr, mux))
}
