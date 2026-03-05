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

func HandlerCreateUser(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	RespondWithJSON(w, 201, DBUserToUser(user))
}

func HandlerGetUserByAPI(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, 200, DBUserToUser(user))
}

func HandlerGetPostsForUser(apiCfg *ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Failed to get posts: %v", err))
		return
	}

	RespondWithJSON(w, 200, DBPostsToPosts(posts))

}
