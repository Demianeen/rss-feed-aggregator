package main

import (
	"time"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	Id        uuid.UUID `json:"id"`
}

func dbUserToUser(user database.User) User {
	return User{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func dbFeedToFeed(feed database.Feed) Feed {
	return Feed{
		Id:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = dbFeedToFeed(dbFeed)
	}
	return feeds
}

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func dbFeedFollowToFeedFolow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		Id:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserId:    dbFeedFollow.UserID,
		FeedId:    dbFeedFollow.FeedID,
	}
}

func dbFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, len(dbFeedFollows))
	for i, dbFeedFollow := range dbFeedFollows {
		feedFollows[i] = dbFeedFollowToFeedFolow(dbFeedFollow)
	}
	return feedFollows
}

type Post struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedId      uuid.UUID `json:"feed_id"`
}

func dbPostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		Id:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		FeedId:      dbPost.FeedID,
	}
}

func dbPostsToPosts(dbPosts []database.Post) []Post {
	feedFollows := make([]Post, len(dbPosts))
	for i, dbFeedFollow := range dbPosts {
		feedFollows[i] = dbPostToPost(dbFeedFollow)
	}
	return feedFollows
}
