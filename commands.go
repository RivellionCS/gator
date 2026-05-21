package main

import (
	"errors"

	"github.com/RivellionCS/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	Name string
	Arguments []string
}

type commands struct {
	command_list map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	fn, ok := c.command_list[cmd.Name]
	if ok {
		fn(s, cmd)
		return nil
	}
	return errors.New("could not run command")
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.command_list[name] = f
}