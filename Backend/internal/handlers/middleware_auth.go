package handlers

import (
	"fmt"
	"net/http"

	"github.com/oneelabed/RSSAggregator/internal/auth"
	. "github.com/oneelabed/RSSAggregator/internal/config"
	"github.com/oneelabed/RSSAggregator/internal/database"
)

type authedHandler func(*ApiConfig, http.ResponseWriter, *http.Request, database.User)

func MiddlewareAuth(apiCfg *ApiConfig, handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPI(r.Header)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPI(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(apiCfg, w, r, user)
	}
}
