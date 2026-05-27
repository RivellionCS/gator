package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RivellionCS/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	userName := s.cfg.CurrentUserName
	user, err  := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	currentTime := time.Now()
	params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: name,
		Url: url,
		UserID: user.ID,

	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}
	fmt.Printf("Printing feed:\n%v\n", feed)
	return nil
}