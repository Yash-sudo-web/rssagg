package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yash-sudo-web/rssagg/internal/database"
)

func (apiConfig *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errJSON(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		errJSON(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	responseJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiConfig *apiConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {

	feed, err := apiConfig.DB.GetFeeds(r.Context())

	if err != nil {
		errJSON(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	responseJSON(w, http.StatusOK, databaseFeedsToFeeds(feed))
}
