package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Demianeen/rss-feed-aggregator/internal/database"
	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

func (cfg *apiConfig) startScraping(concurrentScrapers int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrentScrapers, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feedToScrape, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(concurrentScrapers))
		if err != nil {
			log.Printf("error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feedToScrape {
			wg.Add(1)

			go cfg.scrapeFeed(wg, feed)
		}
		wg.Wait()
	}
}

func (cfg *apiConfig) scrapeFeed(wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	// we always mark feed as fetched
	_, err := cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := getFeedFromUrl(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
	}

	newPostsCount := 0
	for _, item := range rssFeed.Channel.Item {
		if !postExists(cfg.DB, item.Link) {
			newPostsCount++

			id, created_at, updated_at := utils.NewDbEntry()

			description := sql.NullString{}
			if item.Description != "" {
				description.String = item.Description
				description.Valid = true
			}

			publishedAt, err := parseRssDate(item.PubDate)
			if err != nil {
				log.Printf("couldn't parse date: %v", err)
				continue
			}

			_, err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
				ID:          id,
				CreatedAt:   created_at,
				UpdatedAt:   updated_at,
				Title:       item.Title,
				Description: description,
				PublishedAt: publishedAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
			if err != nil {
				log.Printf("couldn't create post: %v", err)
			}
		}
	}
	log.Printf("Feed %s collected, %v new posts found", feed.Name, newPostsCount)
}

func parseRssDate(dateStr string) (time.Time, error) {
	layouts := []string{
		time.RFC1123,                      // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,                     // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC3339,                      // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,                  // "2006-01-02T15:04:05.999999999Z07:00"
		"Mon, 02 Jan 2006 15:04:05 -0700", // Some feeds use this format
	}

	for _, layout := range layouts {
		time, err := time.Parse(layout, dateStr)
		if err == nil {
			return time, nil
		}
	}

	return time.Time{}, fmt.Errorf("couldn't parse the date %s", dateStr)
}

func postExists(db *database.Queries, url string) bool {
	_, err := db.GetPostByUrl(context.Background(), url)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("error checking if post exists: %v", err)
		return false
	}
	return true
}
