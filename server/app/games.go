package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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

// GetGames gets all the messages. This should probably be on a per...game basis
func GetGames(w http.ResponseWriter, r *http.Request) {
	games := Games{}

	err := queryGames(&games)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(games)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}

func queryGames(games *Games) error {
	rows, err := DB.Query(`SELECT * FROM games;`)
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
