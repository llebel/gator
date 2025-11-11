package main

import (
	"fmt"

	"github.com/llebel/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires exactly one argument: username")
	}
	username := cmd.Args[0]
	config.SetUser(username)
	fmt.Printf("User set to '%s'\n", username)
	return nil
}
