package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/RivellionCS/gator/internal/config"
	"github.com/RivellionCS/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	newConfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", newConfig.DbURL)
	if err != nil {
		fmt.Printf("Error connecting to db: %v\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	newState := state{
		db: dbQueries,
		cfg: &newConfig,
	}
	newCommands := commands{
		command_list: map[string]func(*state, command) error{},
	}
	newCommands.register("login", handlerLogin)
	newCommands.register("register", handlerRegister)
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("arguments can't be less than 2")
		os.Exit(1)
	}
	commandName := arguments[1]
	commandSlice := arguments[2:]
	newCommand := command{
		name: commandName,
		arguments: commandSlice,
	}
	err = newCommands.run(&newState, newCommand)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}