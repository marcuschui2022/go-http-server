package main

import (
	"encoding/json"
	"example.com/marcus/go-http-server/internal/auth"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}
	type resp struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	var p parameters
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't decode parameters", err)
		return
	}

	if p.ExpiresInSeconds == 0 {
		p.ExpiresInSeconds = 3600
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), p.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(p.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expireIn := time.Duration(p.ExpiresInSeconds) * time.Second

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expireIn)

	respondWithJSON(w, http.StatusOK, resp{
		User{
			ID:    user.ID,
			Email: user.Email,
			Token: token,
		}})
}
