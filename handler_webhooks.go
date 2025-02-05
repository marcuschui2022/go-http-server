package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), params.Data.UserID)

	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	//respondWithError(w, http.StatusUnauthorized, "", nil)
		//	w.WriteHeader(http.StatusNoContent)
		//	return
		//}
		//respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
