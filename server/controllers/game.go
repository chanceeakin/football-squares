package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	game "football-squares/server/models/games"
	response "football-squares/server/response"
	routes "football-squares/server/routes"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"log"
	"net/http"
)

// GameRoutes is the declaration for all routes
func GameRoutes() []routes.Route {
	gameRoutes := make([]routes.Route, 3)
	gameRoutes = append(gameRoutes,
		routes.Route{
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
		},
		routes.Route{
			Name:        "Archive Game",
			Path:        "/game/archive",
			HandlerFunc: archiveGame,
			Method:      "PUT",
		},
	)
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
	v := validator.New()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	err = v.Struct(input)
	defer r.Body.Close()
	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	val, err1 := game.QueryGame(&input)
	if err1 != nil {
		response.SendError(w, err1, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, val)
}

func postGame(w http.ResponseWriter, r *http.Request) {
	var err error
	var input game.PostInput
	v := validator.New()
	var out common.ID
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	defer r.Body.Close()

	err = v.Struct(input)

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

func archiveGame(w http.ResponseWriter, r *http.Request) {
	var err error
	v := validator.New()
	input := new(common.ID)
	var out common.Success

	err = json.NewDecoder(r.Body).Decode(input)
	defer r.Body.Close()

	err = v.Struct(input)

	switch {
	case err == io.EOF:
	case err != nil:
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	out, err = game.ArchiveGame(input)
	if err != nil {
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}
	response.SendJSON(w, out)

}
