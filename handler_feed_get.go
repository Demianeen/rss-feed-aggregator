package main

import (
	"log"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Printf("Error retrieving feeds: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve feeds")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, dbFeedsToFeeds(feeds))
}
