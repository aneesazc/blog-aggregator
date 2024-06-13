package main

import (
	"net/http"

	"github.com/aneesazc/blog-aggregator/internal/database"
)

func(cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(w, http.StatusOK, user)
}