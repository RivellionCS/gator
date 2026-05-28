package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RivellionCS/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("The follow command requires a url argument")
	}

	url := cmd.arguments[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v", err)
	}

	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error getting user by name: %v", err)
	}

	currentTime := time.Now()
	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Printf("Feed Name: %v\n", feedFollow.FeedName)
	fmt.Printf("Current User: %v\n", feedFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting follows for user: %v", err)
	}

	fmt.Printf("Printing feeds for user: %v\n", username)
	for _, follow := range follows {
		fmt.Printf("Feed Name: %v\n", follow.FeedName)
	}
	return nil
}