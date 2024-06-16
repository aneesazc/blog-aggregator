package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    }

    type response struct {
        Feed        database.Feed       `json:"feed"`
        FeedFollow  database.FeedFollow `json:"feed_follow"`
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
        Name:      params.Name,
        Url:       params.URL,
        UserID:    user.ID,
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
    })

    if err != nil {
        log.Printf("Error creating feed: %v", err)
        respondWithError(w, http.StatusBadRequest, "Failed to create feed")
        return
    }

    log.Println("Feed created successfully, proceeding to create feed follow")

    feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID:    user.ID,
        FeedID:    feed.ID,
    })
    if err != nil {
        log.Printf("Error creating feed follow: %v", err)
        respondWithError(w, http.StatusBadRequest, "Failed to create feed follow")
        return
    }

    log.Println("Feed follow created successfully")

    respondWithJSON(w, http.StatusCreated, response{feed, feedFollow})
}

