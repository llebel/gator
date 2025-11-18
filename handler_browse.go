package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/llebel/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("browse command takes at most one argument: limit")
	}
	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:   user.Name,
		Limit:  int32(limit),
		Offset: 0,
	})
	if err != nil {
		return fmt.Errorf("error retrieving posts for user '%s': %v", user.Name, err)
	}

	for _, post := range posts {
		publishedAt := "N/A"
		if post.PublishedAt.Valid {
			publishedAt = post.PublishedAt.Time.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("Title: %s\nURL: %s\nPublished At: %s\n\n", post.Title, post.Url, publishedAt)
	}

	return nil
}
