package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command does not take any arguments")

	}
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting users: %v", err)
	}

	fmt.Println("All users have been reset")
	return nil
}
