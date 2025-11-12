package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("users command does not take any arguments")

	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("error retrieving users: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Registered users:")
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
