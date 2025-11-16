package main

import (
	"context"
	"fmt"

	"github.com/llebel/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires exactly one argument: username")
	}
	username := cmd.Args[0]

	// Check if user actually exists
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user '%s' does not exist", username)
	}

	config.SetUser(username)
	fmt.Printf("User set to '%s'\n", username)
	return nil
}
