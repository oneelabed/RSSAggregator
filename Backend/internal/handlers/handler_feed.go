package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	. "github.com/oneelabed/RSSAggregator/internal/config"
	"github.com/oneelabed/RSSAggregator/internal/database"
)

func HandlerCreateFeed(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	RespondWithJSON(w, 201, DBFeedToFeed(feed))
}

func HandlerGetFeeds(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	RespondWithJSON(w, 201, DBFeedsToFeeds(feeds))
}
