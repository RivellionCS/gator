package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	s.Config.SetUser(cmd.Arguments[0])
	fmt.Println("the user has been set")
	return nil
}