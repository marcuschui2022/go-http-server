package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

// Chirpy start!
func main() {
	const filepathRoot = "."
	const apiPrefix = "/api/"
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET "+apiPrefix+"healthz", handlerReadiness)
	mux.HandleFunc("GET "+apiPrefix+"metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST "+apiPrefix+"reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
