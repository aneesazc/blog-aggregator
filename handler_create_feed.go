package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	/*
	{
		"name": "The Boot.dev Blog",
		"url": "https://blog.boot.dev/index.xml"
		}
	*/
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if params.Name == "" || params.URL == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}