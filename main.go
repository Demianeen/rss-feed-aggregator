package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the .env environment")
	}
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error opening connection with database. Is postgres running and CONN env is set?")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	go apiCfg.startScraping(10, time.Minute)

	v1Mux := http.NewServeMux()

	v1Mux.HandleFunc("POST /v1/users", apiCfg.handlerCreateUser)
	v1Mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Mux.HandleFunc("GET /v1/posts", apiCfg.middlewareAuth(apiCfg.handlerGetUserPosts))

	v1Mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Mux.HandleFunc("GET /v1/feeds", apiCfg.handlerGetFeeds)

	v1Mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetUserFeedFollows))
	v1Mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	v1Mux.HandleFunc("GET /v1/readiness", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Status string `json:"status"`
		}
		utils.RespondWithJson(w, 200, response{
			Status: "ok",
		})
	})

	v1Mux.HandleFunc("GET /v1/err", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithError(w, 500, "Internal server error")
	})

	server := http.Server{
		Addr:    ":" + port,
		Handler: v1Mux,
	}
	log.Printf("The server is running on port %s", port)
	log.Fatal(server.ListenAndServe())
}
