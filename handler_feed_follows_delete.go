package main

import (
	"fmt"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := r.PathValue("feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Malformed feed follow id")
		return
	}

	_, err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowId,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete feed: %v", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, struct{}{})
}
