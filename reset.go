package main

import (
	"fmt"
	"net/http"
)

func handlerReset(apiCfg *apiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		apiCfg.fileserverHits.Store(0)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		hits := apiCfg.fileserverHits.Load()
		w.Write([]byte(fmt.Sprintf("Hits: %d", hits)))
	}
}
