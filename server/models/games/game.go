package game

import (
	"time"

	db "github.com/chanceeakin/football-squares/server/db"
)

// Game is a data struct for a given game
type Game struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	BegunAt    *time.Time `json:"begun_at"`
	FinishedAt *time.Time `json:"finished_at"`
}

// Games is a slice of game.
type Games struct {
	Games []Game
}

// GetInput is the struct for a single game id.
type GetInput struct {
	ID string `json:"id"`
}

// QueryGames for a series of games
func QueryGames(games *Games) error {
	rows, err := db.DB.Query(`SELECT * FROM games;`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		game := Game{}
		err = rows.Scan(
			&game.ID,
			&game.Title,
			&game.CreatedAt,
			&game.UpdatedAt,
			&game.BegunAt,
			&game.FinishedAt,
		)
		if err != nil {
			return err
		}
		games.Games = append(games.Games, game)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// QueryGame for a series of games
func QueryGame(input *GetInput) (Game, error) {
	returnGame := Game{}
	row := db.DB.QueryRow(`SELECT * FROM games where id=$1;`, &input.ID)
	err := row.Scan(
		&returnGame.ID,
		&returnGame.Title,
		&returnGame.CreatedAt,
		&returnGame.UpdatedAt,
		&returnGame.BegunAt,
		&returnGame.FinishedAt,
	)
	if err != nil {
		return returnGame, err
	}
	return returnGame, nil
}
