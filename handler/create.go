package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"game/entity"
	"log"
	"os"
	"strconv"
	"time"
)

func CreateGame(db *sql.DB) {
	var game entity.Game

	fmt.Print("Enter the name: ")
	nameScanner := bufio.NewScanner(os.Stdin)
	for nameScanner.Scan() {
		game.Name = nameScanner.Text()
		break
	}

	fmt.Print("Enter the description: ")
	descScanner := bufio.NewScanner(os.Stdin)
	for descScanner.Scan() {
		game.Description = descScanner.Text()
		break
	}

	game.PublishedDate = time.Now()

	for {
		if game.Rating > 0 {
			break
		}

		fmt.Print("Enter the rating: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil || input < 0 || input > 5.0 {
				fmt.Println("\nInvalid rating. Minimum score is 0.1 and maximum score is 5")
				continue
			}

			game.Rating = input
			break
		}
	}

	_, err := db.Exec("INSERT INTO game (name, description, published_date, rating) VALUES (?, ?, ?, ?)", game.Name, game.Description, game.PublishedDate, game.Rating)
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}

	fmt.Printf("Game: %s created successfully.\n", game.Name)
}
