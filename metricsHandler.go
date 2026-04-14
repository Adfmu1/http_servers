package main

import (
	"net/http"
	"fmt"
)

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