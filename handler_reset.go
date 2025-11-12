package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command does not take any arguments")

	}
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Printf("error resetting users: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All users have been reset")
	return nil
}
