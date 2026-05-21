package main

import (
	"errors"

	"github.com/RivellionCS/gator/internal/config"
)

type state struct {
	config *config.Config
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
		fn(s, cmd)
		return nil
	}
	return errors.New("could not run command")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.command_list[name] = f
}