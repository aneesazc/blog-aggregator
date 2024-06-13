package main

import "net/http"

func(cfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request){
	feed, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feed)
}