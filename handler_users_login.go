package main

import (
	"encoding/json"
	"example.com/marcus/go-http-server/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	var p parameters
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "wrong id/password", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), p.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "", err)
		return
	}

	err = auth.CheckPasswordHash(p.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}
