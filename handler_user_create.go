package main

import (
	"log"
	"net/http"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (c *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	var params parameters
	if !utils.DecodeJsonBody(w, r, &params) {
		return
	}

	id, createdAt, updatedAt := utils.NewDbEntry()
	arg := database.CreateUserParams{
		Name:      params.Name,
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	user, err := c.DB.CreateUser(r.Context(), arg)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, dbUserToUser(user))
}
