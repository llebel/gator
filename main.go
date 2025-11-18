package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/llebel/gator/internal/config"
	"github.com/llebel/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg := config.Read()
	s := state{
		cfg: &cfg,
	}

	// Register commands
	commands := commands{commandMap: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	// Register database
	db, err := sql.Open("postgres", s.cfg.DbURL)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	s.db = database.New(db)

	// Parse command line
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{Name: cmdName, Args: cmdArgs}

	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
