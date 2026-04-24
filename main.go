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
	SecretKey		string
}

var apiConf apiConfig

func main() {
	godotenv.Load()
	// import db
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("an error occured while opening database")
		return
	}
	dbQueries := database.New(dbConn)

	apiConf.Database = dbQueries
	apiConf.Platform = platform
	apiConf.SecretKey = secret

	// create a multiplexer for a server
	mux := http.NewServeMux()
	// create server that uses created mux
	serv := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}

	// add basic handler at a root
	mux.Handle("/app/", http.StripPrefix("/app/", apiConf.middlewareMetricsInc(http.FileServer(http.Dir(".")))))

	mux.HandleFunc("GET /api/healthz", handleReadinessEndpoint)
	mux.HandleFunc("GET /api/chirps", handleGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirp_id}", handleGetChirp)

	mux.HandleFunc("POST /api/chirps", handlePostChirps)
	mux.HandleFunc("POST /api/users", handleRegisterUser)
	mux.HandleFunc("POST /api/login", handleLoginUser)
	mux.HandleFunc("POST /api/refresh", handleRefresh)
	mux.HandleFunc("POST /api/revoke", handleRevoke)

	mux.HandleFunc("PUT /api/users", editUserDataHandler)

	mux.HandleFunc("GET /admin/metrics", apiConf.handleMetricsEndpoint)
	mux.HandleFunc("POST /admin/reset", apiConf.handleResetEndpoint)

	serv.ListenAndServe()
}
