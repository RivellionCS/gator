package main

import (
	"errors"

	"github.com/RivellionCS/gator/internal/config"
	"github.com/RivellionCS/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	command_list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.command_list[cmd.name]
	if ok {
		return fn(s, cmd)
	}
	return errors.New("could not run command")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.command_list[name] = f
}