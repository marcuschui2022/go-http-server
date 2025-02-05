package main

import (
	"example.com/marcus/go-http-server/internal/auth"
	"example.com/marcus/go-http-server/internal/database"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *apiConfig) handlerDeleteChirpByID(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get token", err)
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
	}

	idString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirpsByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Not found chirpID", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "status forbidden", err)
		return
	}

	err = cfg.db.DeleteChirpsByID(r.Context(), database.DeleteChirpsByIDParams{
		ID:     chirpID,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
