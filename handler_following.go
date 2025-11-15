package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command does not take any arguments")

	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Printf("error retrieving following feeds: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Following feeds:")
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}
