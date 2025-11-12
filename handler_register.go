package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/llebel/gator/internal/config"
	"github.com/llebel/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("register command requires exactly one argument: username")
	}
	username := cmd.Args[0]

	// Check if user already exists
	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		fmt.Printf("user '%s' already exists\n", username)
		os.Exit(1)
	}

	// Creating a new user
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		Name:      username,
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt: sql.NullTime{Time: time.Now()},
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	// Set as current user
	config.SetUser(username)
	fmt.Printf("Registered user '%s'\n", username)
	fmt.Printf("%v\n", newUser)
	return nil
}
