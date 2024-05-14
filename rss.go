package main

import (
	"encoding/xml"

	"github.com/Demianeen/rss-feed-aggregator/internal/utils"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func parseRssFeed(data []byte) (RssFeed, error) {
	rssFeed := RssFeed{}
	err := xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RssFeed{}, err
	}

	return rssFeed, nil
}

func getFeedFromUrl(url string) (RssFeed, error) {
	data, err := utils.FetchDataFromUrl(url)
	if err != nil {
		return RssFeed{}, err
	}

	rssFeed, err := parseRssFeed(data)
	if err != nil {
		return RssFeed{}, err
	}

	return rssFeed, nil
}
