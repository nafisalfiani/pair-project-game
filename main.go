package main

import (
	"database/sql"
	"fmt"
	"game/handler"
	"log"
	"os"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/game_library")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\nWelcome to Game Library! Please select from given command")

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "-------------\t | -------------------------\t")
		fmt.Fprintln(w, "COMMAND\t | DESCRIPTION\t")
		fmt.Fprintln(w, "-------------\t | -------------------------\t")

		fmt.Fprintln(w, "[list-games]\t | see list of games")
		fmt.Fprintln(w, "[create-game]\t | create new game entry")
		fmt.Fprintln(w, "[update-game]\t | update existing game entry")
		fmt.Fprintln(w, "[delete-game]\t | delete existing game entry")
		w.Flush()

		fmt.Println("Press anything to back or exit")
		fmt.Println()

		var choice string
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Exiting...")
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		switch choice {
		case "create-game":
			handler.CreateGame(db)
		case "list-games":
			handler.GetAllGames(db)
		case "update-game":
			handler.UpdateGame(db)
		case "delete-game":
			handler.DeleteGame(db)
		case "":
			fmt.Println("Exiting...")
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Please enter valid command.")
			fmt.Println()
		}
	}
}
