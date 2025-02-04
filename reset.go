package main

import (
	"net/http"
	"os"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")
	if platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err := cfg.db.ResetUser(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't reset user", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	//cfg.fileserverHits.Store(0)
	//w.WriteHeader(http.StatusOK)
	//if _, err := w.Write([]byte("Hits reset to 0")); err != nil {
	//	log.Printf("Error writing response: %v", err)
	//}
}
