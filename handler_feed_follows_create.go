package main

import (
	"log"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	var params parameters
	if !utils.DecodeJsonBody(w, r, &params) {
		return
	}

	id, createdAt, updatedAt := utils.NewDbEntry()
	newFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    params.FeedId,
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		log.Printf("Error creating feed: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, dbFeedFollowToFeedFolow(newFeedFollow))
}
