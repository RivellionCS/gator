package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/RivellionCS/gator/internal/database"
	"github.com/google/uuid"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the agg command expects 1 time argument")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error parsing time: %v", err)
	}
	ticker := time.NewTicker(timeBetweenReqs)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			log.Printf("error scraping feeds: %v", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed: %v", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %v", err)
	}
	for _, item := range rssFeed.Channel.Item {
		publishedTime := sql.NullTime{}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			publishedTime = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}

		postDescription := sql.NullString{
			String: item.Description,
			Valid: true,
		}

		currentTime := time.Now()
		params := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
			Title: item.Title,
			Url: item.Link,
			Description: postDescription,
			PublishedAt: publishedTime,
			FeedID: feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	return nil
}