package game

import (
	"log"
	"time"

	common "football-squares/server/common"
	db "football-squares/server/db"
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

// PostInput is the input for a new game
type PostInput struct {
	Title string `json:"title"`
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
func QueryGame(input *common.ID) (Game, error) {
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

// PostGame savess a game record in the database
func PostGame(input *PostInput) (common.ID, error) {
	var err error
	insertStatement := `
	INSERT INTO games (title)
	VALUES ($1)
	RETURNING id;`
	out := common.ID{}
	err = db.DB.QueryRow(insertStatement, &input.Title).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}
