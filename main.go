package main

import (
	"fmt"

	"github.com/llebel/gator/internal/config"
)

func main() {
	cfg := config.Read()
	config.SetUser("llebel")
	cfg = config.Read()
	fmt.Println("Current User:", cfg.CurrentUserName)
	fmt.Println("DbURL:", cfg.DbURL)
}
