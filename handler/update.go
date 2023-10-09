package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"game/entity"
	"log"
	"os"
	"strconv"
)

func UpdateGame(db *sql.DB) {
	var gameID int
	fmt.Print("Enter the ID of the game you want to update: ")
	_, err := fmt.Scanln(&gameID)
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}

	var game entity.Game

	fmt.Print("Enter the new name (leave empty to keep existing): ")
	nameScanner := bufio.NewScanner(os.Stdin)
	for nameScanner.Scan() {
		if nameScanner.Text() != "" {
			game.Name = nameScanner.Text()
		}

		break
	}

	fmt.Print("Enter the new description (leave empty to keep existing): ")
	descScanner := bufio.NewScanner(os.Stdin)
	for descScanner.Scan() {
		if descScanner.Text() != "" {
			game.Description = descScanner.Text()
		}

		break
	}

	fmt.Print("Enter the new rating (leave empty to keep existing): ")
	ratingScanner := bufio.NewScanner(os.Stdin)
	for ratingScanner.Scan() {
		if ratingScanner.Text() == "" {
			break
		}

		input, err := strconv.ParseFloat(ratingScanner.Text(), 64)
		if err != nil || input < 0 || input > 5.0 {
			fmt.Println("\nInvalid rating. Minimum score is 0.1 and maximum score is 5. Skipping rating update")
			break
		}

		game.Rating = input
		break
	}

	if game.Name == "" && game.Description == "" && game.Rating == 0 {
		fmt.Println("No updates provided. Game details remain unchanged.")
		return
	}

	_, err = db.Exec("UPDATE game SET name=COALESCE(?, name), description=COALESCE(?, description), rating=COALESCE(?, rating) WHERE game_id=?", game.Name, game.Description, game.Rating, gameID)
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}

	fmt.Println("Game updated successfully.")
}
