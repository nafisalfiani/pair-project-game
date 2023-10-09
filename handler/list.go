package handler

import (
	"database/sql"
	"fmt"
	"game/entity"
	"log"
	"time"
)

func GetAllGames(db *sql.DB) {
	rows, err := db.Query("SELECT game_id, name, description, published_date, rating FROM game")
	if err != nil {
		log.Println("Sorry, game library is not available right now. Please come back later")
		return
	}
	defer rows.Close()

	var games []entity.Game
	for rows.Next() {
		var publishedDate string
		var game entity.Game
		err := rows.Scan(&game.GameID, &game.Name, &game.Description, &publishedDate, &game.Rating)
		if err != nil {
			log.Println("Sorry, game library is not available right now. Please come back later")
			return
		}

		game.PublishedDate, err = time.Parse("2006-01-02", publishedDate)
		if err != nil {
			log.Println("Sorry, game library is not available right now. Please come back later")
			return
		}

		games = append(games, game)
	}

	fmt.Println("Game Library List:")
	for _, game := range games {
		fmt.Printf("[%d] Name: %s, Description: %s, Published Date: %s, Rating: %.1f\n",
			game.GameID, game.Name, game.Description, game.PublishedDate.Format("2006-01-02"), game.Rating)
	}
}
