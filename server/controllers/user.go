package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	user "football-squares/server/models/user"
	response "football-squares/server/response"
	routes "football-squares/server/routes"
	"gopkg.in/go-playground/validator.v9"
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
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, userMap)
}

func getUser(w http.ResponseWriter, r *http.Request) {
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

	val, err1 := user.QueryUser(&input)
	if err1 != nil {
		response.SendError(w, err1, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, val)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var userInput user.Input
	var out common.ID
	v := validator.New()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userInput)
	defer r.Body.Close()
	err = v.Struct(userInput)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	out, err = user.InsertUser(&userInput)
	if err != nil {
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}
	response.SendJSON(w, out)
}
