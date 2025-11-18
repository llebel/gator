package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg command takes exactly one argument: time_between_reqs")
	}

	// Parse time between requests
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}
	fmt.Printf("Aggregating feeds with %v between requests...\n", time_between_reqs)

	// Fetch the feed
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
