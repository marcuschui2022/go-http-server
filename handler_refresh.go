package main

import (
	"example.com/marcus/go-http-server/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Token string `json:"token"`
	}

	requestToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), requestToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get user for refresh token", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't access jwt token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, resp{
		Token: token,
	})

}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	requestToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}

	err = cfg.db.RevokeToken(r.Context(), requestToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
