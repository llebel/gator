package main

import (
	"context"
	"fmt"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("agg command does not take any arguments")
	}

	rss, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %v\n", rss)
	return nil
}
