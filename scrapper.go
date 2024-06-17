package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/aneesazc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

// The function initiates multiple goroutines to scrape feeds concurrently.
func startScraping(db *database.Queries, conc int, timeBwRequest time.Duration){
	log.Printf("Starting scraping with %d goroutines every %s duration\n", conc, timeBwRequest)
	ticker := time.NewTicker(timeBwRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(conc))
		if err != nil {
			log.Printf("Error getting feeds to fetch: %v\n", err)
			continue
		}
		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, feed, &wg)
		}
		wg.Wait()
	}

}

// The function is called for each feed, marking it as fetched and processing the feed items.
func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup){
	defer wg.Done()
	// Scrape the feed
	log.Printf("Scraping feed: %s\n", feed.Url)
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error scraping feed: %v\n", err)
		return
	}	

	for _, item := range rssFeed.Channel.Item {
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			if strings.Contains(err.Error(), "parsing time") {
				continue
			}
			log.Printf("Error parsing time: %v\n", err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID: feed.ID,
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: pubAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error creating post: %v\n", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}