package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/llebel/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func unescapeHTML(rss *RSSFeed) {
	// Implementation for unescaping HTML entities in RSS feed fields
	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Implementation for fetching and parsing the RSS feed
	client := &http.Client{}

	// Creating HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator/1.0")

	// Making the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parsing the RSS feed from response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rss RSSFeed
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}

	// Unescape HTML entities in the feed
	unescapeHTML(&rss)

	return &rss, nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// No feed to scrape
	if nextFeed.ID == uuid.Nil {
		return nil
	}

	// Mark feed as fetched
	err = s.db.MarkFeedFetchedbyID(context.Background(), database.MarkFeedFetchedbyIDParams{
		ID:            nextFeed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	// Scrape the feed
	fmt.Printf("Scraping feed: %s (%s)\n", nextFeed.Name, nextFeed.Url)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	// Process each item in the feed
	for _, item := range rssFeed.Channel.Item {
		// Parse publication date
		var pubDate sql.NullTime
		if item.PubDate != "" {
			parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				parsedTime, err = time.Parse(time.RFC1123, item.PubDate)
				if err != nil {
					pubDate = sql.NullTime{Valid: false}
				} else {
					pubDate = sql.NullTime{Time: parsedTime, Valid: true}
				}
			} else {
				pubDate = sql.NullTime{Time: parsedTime, Valid: true}
			}
		} else {
			pubDate = sql.NullTime{Valid: false}
		}

		// Create post in the database
		fmt.Printf("Saving post %s (%s)\n", item.Title, item.Link)
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubDate,
			FeedID:      uuid.NullUUID{UUID: nextFeed.ID, Valid: true},
		})
		if err != nil {
			// Ignore unique constraint violation errors
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				continue
			}
			return err
		}
	}

	return nil
}
