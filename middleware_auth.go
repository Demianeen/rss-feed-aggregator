package main

import (
	"fmt"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/auth"
	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
