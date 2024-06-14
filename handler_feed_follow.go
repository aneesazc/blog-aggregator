package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if params.FeedID == "" {
		respondWithError(w, http.StatusBadRequest, "Feed ID is required")
		return
	}

	feed_follow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: uuid.MustParse(params.FeedID),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to create feed follow")
		return
	}

	respondWithJSON(w, http.StatusCreated, feed_follow)
}