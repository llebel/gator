package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/llebel/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed command requires exactly two arguments: feed_name and feed_url")
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// Getting current user
	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("error retrieving current user '%s': %v\n", username, err)
		os.Exit(1)
	}

	// Creating a new feed
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	// Following the created feed
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: newFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following the created feed: %v", err)
	}

	// Print created feed
	fmt.Printf("%v\n", newFeed)
	return nil
}
