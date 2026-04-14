package main

import _ "github.com/lib/pq"

import (
	"github.com/joho/godotenv"
	"github.com/Adfmu1/http_servers/internal/database"
	"database/sql"
	"net/http"
	"os"
)


func main() {
	godotenv.Load()
	// import db
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	// create a multiplexer for a server
	mux := http.NewServeMux()
	// create server that uses created mux
	serv := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}
	apiConf := &apiConfig{}
	// add basic handler at a root
	mux.Handle("/app/", http.StripPrefix("/app/", apiConf.middlewareMetricsInc(http.FileServer(http.Dir(".")))))

	mux.HandleFunc("GET /api/healthz", handleReadinessEndpoint)
	mux.HandleFunc("POST /api/validate_chirp", handlePostChirps)

	mux.HandleFunc("GET /admin/metrics", apiConf.handleMetricsEndpoint)
	mux.HandleFunc("POST /admin/reset", apiConf.handleResetEndpoint)

	serv.ListenAndServe()
}
