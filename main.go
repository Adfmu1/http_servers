package main

import (
	"net/http"
	"sync/atomic"
	"encoding/json"
	"fmt"
)


// struct to hold number of requests to the server
type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
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
	// readiness endpoint handler
	mux.HandleFunc("GET /api/healthz", handleReadinessEndpoint)
	mux.HandleFunc("POST /api/validate_chirp", handlePostChirps)

	mux.HandleFunc("GET /admin/metrics", apiConf.handleMetricsEndpoint)
	mux.HandleFunc("POST /admin/reset", apiConf.handleResetEndpoint)
	// start the server
	serv.ListenAndServe()
}

// handler function for readiness endpoint
func handleReadinessEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}

// handler for accepting Chirps
func handlePostChirps(rw http.ResponseWriter, req *http.Request) {
	// decode the data from request
	type parameters struct {
		Body	string	`json:"body"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	// ======== CHECK REQUEST DATA ========
	// any error occurs
	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}
	// length of chirp too long
	if len(params.Body) > 140 {
		const errMsg = "Chirp is too long"
		respondWithError(rw, 400, errMsg)
		return
	}
	// response ok
	type response struct {
		Body	bool	`json:"valid"`
	}
	resp := response{
		Body:	true,
	}
	respondWithJson(rw, 200, resp)
}

// ======== METHODS for apiConfig ========
// handler method for metrics
func (cfg *apiConfig) handleMetricsEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(200)
	adminPageContent := fmt.Sprintf(`<html>
	  <body>
	    <h1>Welcome, Chirpy Admin</h1>
	    <p>Chirpy has been visited %d times!</p>
	  </body>
	</html>`, cfg.fileserverHits.Load())
	rw.Write([]byte(adminPageContent))
} 

// handler method for reseting metrics
func (cfg *apiConfig) handleResetEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}

// middleware method for counting requests
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// ======== HELPER FUNCTIONS =========
func respondWithError(rw http.ResponseWriter, code int, msg string) {
	type errorStruct struct {
		ErrMsg	string	`json:"error"`
	}
	errDat := errorStruct{
		ErrMsg:	msg,
	}
	respondWithJson(rw, code, errDat)
}

func respondWithJson(rw http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	// any error occurs
	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(data)
}

