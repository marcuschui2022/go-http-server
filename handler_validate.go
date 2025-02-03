package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool   `json:"valid"`
		Body  string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	fmt.Println(params.Body)
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Fields(params.Body)
	for i, word := range words {
		for _, badWord := range badWords {
			if strings.EqualFold(strings.ToLower(word), badWord) {
				words[i] = "****"
			}
		}

	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
		Body:  strings.Join(words, " "),
	})

}
