package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("feeds command does not take any arguments")

	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %v", err)
	}

	fmt.Println("Registered feeds:")
	for _, feed := range feeds {
		// Getting user name for the feed
		// Note: This is not optimal, as it performs one query per feed, we should join instead
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error retrieving user for feed '%s': %v", feed.Name, err)
		}
		fmt.Printf("* %s (%s) created by: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
