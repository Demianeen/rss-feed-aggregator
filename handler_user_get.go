package main

import (
	"fmt"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, err := utils.GetQueryParamInt(r.URL.Query(), "limit", 10)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse limit param: %v", err))
		return
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Couldn't get posts for a user")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, posts)
}
