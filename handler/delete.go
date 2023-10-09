package handler

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteGame(db *sql.DB) {
	var gameID int
	fmt.Print("Enter the ID of the game you want to delete: ")
	_, err := fmt.Scanln(&gameID)
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}

	_, err = db.Exec("DELETE FROM game WHERE game_id=?", gameID)
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}

	fmt.Println("Game deleted successfully.")
}
