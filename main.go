package main

import _ "github.com/lib/pq"

import (
	"github.com/joho/godotenv"
	"github.com/Adfmu1/http_servers/internal/database"
	"sync/atomic"
	"database/sql"
	"net/http"
	"os"
	"fmt"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	Database	   *database.Queries
	Platform		string
}

var apiConf apiConfig

func main() {
	godotenv.Load()
	// import db
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("an error occured while opening database")
		return
	}
	dbQueries := database.New(dbConn)

	// create a multiplexer for a server
	mux := http.NewServeMux()
	// create server that uses created mux
	serv := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}
	apiConf.Database = dbQueries
	apiConf.Platform = platform

	// add basic handler at a root
	mux.Handle("/app/", http.StripPrefix("/app/", apiConf.middlewareMetricsInc(http.FileServer(http.Dir(".")))))

	mux.HandleFunc("GET /api/healthz", handleReadinessEndpoint)
	mux.HandleFunc("POST /api/validate_chirp", handlePostChirps)
	mux.HandleFunc("POST /api/users", handleRegisterUser)

	mux.HandleFunc("GET /admin/metrics", apiConf.handleMetricsEndpoint)
	mux.HandleFunc("POST /admin/reset", apiConf.handleResetEndpoint)

	serv.ListenAndServe()
}
