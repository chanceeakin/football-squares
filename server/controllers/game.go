package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	game "football-squares/server/models/games"
	response "football-squares/server/response"
	routes "football-squares/server/routes"

	"log"
	"net/http"
)

// GameRoutes is the declaration for all routes
func GameRoutes() []routes.Route {
	gameRoutes := make([]routes.Route, 3)
	gameRoutes = append(gameRoutes, routes.Route{
		Name:        "Games",
		Path:        "/games",
		HandlerFunc: getGames,
		Method:      "GET",
	},
		routes.Route{
			Name:        "Game",
			Path:        "/game",
			HandlerFunc: getGame,
			Method:      "GET",
		},
		routes.Route{
			Name:        "Game",
			Path:        "/game",
			HandlerFunc: postGame,
			Method:      "POST",
		})
	return gameRoutes
}

func getGames(w http.ResponseWriter, r *http.Request) {
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
