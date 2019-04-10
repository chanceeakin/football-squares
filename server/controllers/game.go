package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	game "football-squares/server/models/games"
	response "football-squares/server/response"
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

	response.SendJSON(w, gamesArr)
}

// GetGame gets a single game
func getGame(w http.ResponseWriter, r *http.Request) {
	var input common.ID
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

	response.SendJSON(w, val)
}

func postGame(w http.ResponseWriter, r *http.Request) {
	var err error
	var input game.PostInput
	var out common.ID
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)

	if err != nil {
		log.Print(err)
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	out, err = game.PostGame(&input)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}
	response.SendJSON(w, out)

}

// GameHandler is the switch for REST Methods
func GameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getGame(w, r)
	case http.MethodPost:
		postGame(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, `Not Found`, http.StatusNotFound)
	}
}
