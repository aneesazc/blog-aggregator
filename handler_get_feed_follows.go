package main

import (
	"net/http"

	"github.com/aneesazc/blog-aggregator/internal/database"
)

func(cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	feed_follows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to get feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, feed_follows)
}