package main

import (
	"log"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	type response struct {
		Feed       `json:"feed"`
		FeedFollow `json:"feed_follow"`
	}

	var params parameters
	if !utils.DecodeJsonBody(w, r, &params) {
		return
	}

	feedId, createdAt, updatedAt := utils.NewDbEntry()
	newFeed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
		ID:        feedId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		log.Printf("Error creating feed: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	// we subscribe user automatically to the feed it created
	feedFollowId, createdAt, updatedAt := utils.NewDbEntry()
	newFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feedId,
		ID:        feedFollowId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		log.Printf("Error creating feed: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, response{
		Feed:       dbFeedToFeed(newFeed),
		FeedFollow: dbFeedFollowToFeedFolow(newFeedFollow),
	})
}
