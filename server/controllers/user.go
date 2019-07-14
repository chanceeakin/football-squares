package controller

import (
	"encoding/json"
	user "football-squares/server/models/user"
	response "football-squares/server/response"
	routes "football-squares/server/routes"
	"log"
	"net/http"
)

// UserRoutes is the declaration for all routes
func UserRoutes() []routes.Route {
	userRoutes := make([]routes.Route, 3)
	userRoutes = append(userRoutes, routes.Route{
		Name:        "Users",
		Path:        "/users",
		HandlerFunc: getUsers,
		Method:      "GET",
	},
		routes.Route{
			Name:        "Get User",
			Path:        "/user",
			HandlerFunc: getUser,
			Method:      "GET",
		},
		routes.Route{
			Name:        "Post User",
			Path:        "/user",
			HandlerFunc: postUser,
			Method:      "POST",
		})
	return userRoutes
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	userMap := user.Users{}

	err := user.QueryUsers(&userMap)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, userMap)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var input user.GetInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	val, err1 := user.QueryUser(&input)
	if err1 != nil {
		log.Print(err1)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, val)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var userInput user.Input
	var out user.Out
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userInput)

	if err != nil {
		log.Print(err)
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	out, err = user.InsertUser(&userInput)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}
	response.SendJSON(w, out)
}
