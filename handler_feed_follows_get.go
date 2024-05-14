package main

import (
	"log"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (cfg *apiConfig) handlerGetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	userFeedFollows, err := cfg.DB.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		log.Printf("Error getting user feed follows: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't get user's feed follows")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, dbFeedFollowsToFeedFollows(userFeedFollows))
}
