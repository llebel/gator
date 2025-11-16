package main

import (
	"context"
	"fmt"

	"github.com/llebel/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command does not take any arguments")

	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error retrieving following feeds: %v", err)
	}

	fmt.Println("Following feeds:")
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}
