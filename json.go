package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	//if errors.Is(err, sql.ErrNoRows) {
	//	respondWithError(w, http.StatusUnauthorized, "", nil)
	//	return
	//}

	if err != nil {
		log.Printf("Error: %v", err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", err)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}
