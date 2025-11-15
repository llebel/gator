package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/llebel/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command requires exactly one argument: url")
	}
	url := cmd.Args[0]

	// Getting current user
	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("error retrieving current user '%s': %v\n", username, err)
		os.Exit(1)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Printf("error retrieving feed by url '%s': %v\n", url, err)
		os.Exit(1)
	}

	// Follwowing the feed
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}

	// Print followed feed
	fmt.Printf("User '%s' is now following feed '%s'\n", follow.UserName, follow.FeedName)

	return nil
}
