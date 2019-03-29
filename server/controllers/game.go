package controller

import (
	"encoding/json"
	game "github.com/chanceeakin/football-squares/server/models/games"
	"log"
	"net/http"
)

// GetGames gets all the messages. This should probably be on a per...game basis
func GetGames(w http.ResponseWriter, r *http.Request) {
	gamesArr := game.Games{}

	err := game.QueryGames(&gamesArr)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(gamesArr)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}

// GetGame gets a single game
func GetGame(w http.ResponseWriter, r *http.Request) {
	var input game.GetInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	val, err1 := game.QueryGame(&input)
	if err1 != nil {
		log.Print(err1)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(val)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}
