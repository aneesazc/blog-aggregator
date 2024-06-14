package main

import (
	"net/http"
	"strings"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func(cfg *apiConfig) handlerDeleteFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feed_id := ""
	path := r.URL.Path // alternative: r.PathValue("feedFollowID")
    parts := strings.Split(path, "/")
    if len(parts) >= 4 && parts[1] == "v1" && parts[2] == "feed_follows" {
        feedFollowID := parts[3]
		feed_id = feedFollowID
    } else {
        http.NotFound(w, r)
    }

	if feed_id == "" {
		respondWithError(w, http.StatusBadRequest, "Feed ID is required")
		return
	}

	args := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: uuid.MustParse(feed_id),
	}
	err := cfg.DB.DeleteFeedFollow(r.Context(), args)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}