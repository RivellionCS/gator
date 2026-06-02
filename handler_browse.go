package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/RivellionCS/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	browseLimit := 2
	if len(cmd.arguments) == 1 {
		var err error
		browseLimit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("error parsing argument to int: %v\n", err)
		}
	}
	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(browseLimit),
	}
	userPosts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error getting posts for user: %v", err)
	}
	for _, post := range userPosts {
		fmt.Printf("Post URL: %v\n", post.Url)
		fmt.Printf("Post Publish Time: %v\n", post.PublishedAt.Time)
		fmt.Printf("Post Title: %v\n", post.Title)
		fmt.Printf("Post Description: %v\n", post.Description.String)
	}
	return nil
}