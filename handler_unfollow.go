package main

import (
	"context"
	"fmt"

	"github.com/llebel/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("unfollow command requires exactly one argument: url")

	}
	url := cmd.Args[0]

	// Getting feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed by url '%s': %v", url, err)
	}

	// Deleting the feed follow
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	// Print unfollowed feed
	fmt.Printf("User '%s' has unfollowed feed '%s'\n", user.Name, feed.Name)

	return nil
}
