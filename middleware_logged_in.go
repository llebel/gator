package main

import (
	"context"
	"fmt"

	"github.com/llebel/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Getting current user
		username := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("no user is currently logged in, please login first")
		}

		return handler(s, cmd, user)
	}
}
