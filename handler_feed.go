package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RivellionCS/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feeds: %v", err)
	}
	for _,feed := range feeds {
		fmt.Printf("Feed Name: %v\n", feed.FeedsName)
		fmt.Printf("Feed Url: %v\n", feed.FeedsUrl)
		fmt.Printf("Feed User: %v\n", feed.UserName)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("The add feed handler requires the name and the url as arguments")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	userName := s.cfg.CurrentUserName
	user, err  := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	currentTime := time.Now()
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: name,
		Url: url,
		UserID: user.ID,

	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	
	fmt.Println("Printing feed:")
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("Url: %v\n", feed.Url)
	fmt.Printf("UserID: %v\n", feed.UserID)
	return nil
}