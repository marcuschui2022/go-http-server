package main

import (
	"example.com/marcus/go-http-server/internal/database"
	"github.com/google/uuid"
	"net/http"
	"sort"
)

func (cfg *apiConfig) handlerGetChirpByID(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	sortParam := r.URL.Query().Get("sort")
	idString := r.URL.Query().Get("author_id")

	if idString == "" {
		chirps, err := cfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
			return
		}

		respondWithJSON(w, http.StatusOK, sortChirps(chirps, sortParam))
		return
	}

	authorID, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
		return
	}

	chirps, err := cfg.db.GetChirpsByUserID(r.Context(), authorID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	respondWithJSON(w, http.StatusOK, sortChirps(chirps, sortParam))
}

func sortChirps(chirps []database.Chirp, sortParam string) []Chirp {
	var sortedChirps []Chirp
	switch sortParam {
	case "asc":
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.Nanosecond() < chirps[j].CreatedAt.Nanosecond()
		})
		for _, chirp := range chirps {
			sortedChirps = append(sortedChirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
		}
		return sortedChirps
	case "desc":
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.Nanosecond() > chirps[j].CreatedAt.Nanosecond()
		})
		for _, chirp := range chirps {
			sortedChirps = append(sortedChirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
		}
		return sortedChirps
	default:
		for _, chirp := range chirps {
			sortedChirps = append(sortedChirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
		}
		return sortedChirps
	}
}
