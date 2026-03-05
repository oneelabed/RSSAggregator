package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	. "github.com/oneelabed/RSSAggregator/internal/config"
	"github.com/oneelabed/RSSAggregator/internal/database"
)

func HandlerCreateFeedFollow(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	RespondWithJSON(w, 201, DBFollowToFollow(feed_follow))
}

func HandlerGetFeedFollows(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	RespondWithJSON(w, 201, DBFollowsToFollows(follows))
}

func HandlerDeleteFeedFollow(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follow_id_str := chi.URLParam(r, "feedFollowId")
	feed_follow_id, err := uuid.Parse(feed_follow_id_str)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feed_follow_id,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	RespondWithJSON(w, 200, "deleted follow successfully")
}
