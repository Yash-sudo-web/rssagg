package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yash-sudo-web/rssagg/internal/database"
)

func (apiConfig *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		errJSON(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	responseJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseJSON(w, 200, databaseUserToUser(user))
}
