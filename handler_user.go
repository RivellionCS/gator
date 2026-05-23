package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/RivellionCS/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("user %s does not exist: %w", cmd.arguments[0], err)
	}
	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	fmt.Println("the user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the register handler expects a single argument, the name")
	}
	currentTime := time.Now()
	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: cmd.arguments[0],
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't set user %w", err)
	}
	fmt.Printf("A new user has been created:\n%v\n", user)
	return nil
}