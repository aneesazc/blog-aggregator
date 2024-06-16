package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Name string `json:"name"`
    }

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
        return
    }

    if params.Name == "" {
        respondWithError(w, http.StatusBadRequest, "Name is required")
        return
    }

    existingUser, err := cfg.DB.GetUserByName(r.Context(), params.Name)
    if err != nil {
        // Check if the error indicates that the user was not found
        if errors.Is(err, sql.ErrNoRows) {
            // User not found, proceed to create a new user
            user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
                ID:        uuid.New(),
                Name:      params.Name,
                CreatedAt: time.Now().UTC(),
                UpdatedAt: time.Now().UTC(),
            })
            if err != nil {
                respondWithError(w, http.StatusBadRequest, "Failed to create user")
                return
            }
            respondWithJSON(w, http.StatusCreated, user)
            return
        } else {
            // Some other error occurred
            respondWithError(w, http.StatusInternalServerError, "Failed to get user")
            return
        }
    }

    // If no error, it means the user exists
    respondWithJSON(w, http.StatusOK, existingUser)
}


func (cfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User){
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get posts")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)

}