package main

import (
	"fmt"
	"os"

	"github.com/llebel/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg := config.Read()
	s := state{
		cfg: &cfg,
	}

	commands := commands{commandMap: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{Name: cmdName, Args: cmdArgs}

	err := commands.run(&s, cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
