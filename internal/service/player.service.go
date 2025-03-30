package service

import (
	"avatar/internal/db"
	"avatar/internal/models"
	"context"
)

// GetAllPlayers retrieves all players from the database
func GetAllPlayers() ([]models.Player, error) {
	rows, err := db.GetDB().Query(context.Background(), "SELECT id, name, score FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.ID, &player.Name, &player.Score); err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

// UpdatePlayerScore updates the score of a player
func UpdatePlayerScore(id int, score int) error {
	_, err := db.GetDB().Exec(context.Background(), "UPDATE players SET score = score + $1 WHERE id = $2", score, id)
	return err
}
