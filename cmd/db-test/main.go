package main

import (
	"fmt"

	"github.com/psanti93/scotty-g-pizza/db"
)

func main() {
	cfg := db.DefaultConfig()
	db, err := db.Open(cfg)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database pinged and running")
}
