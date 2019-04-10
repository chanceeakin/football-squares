package controller

import (
	"encoding/json"
	user "football-squares/server/models/user"
	response "football-squares/server/response"
	"log"
	"net/http"
)

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

// UsersHandlers is the switch for REST Methods
func UsersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, `Not Found`, http.StatusNotFound)
	}
}

// UserHandlers is the switch for REST Methods
func UserHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		postUser(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, `Not Found`, http.StatusNotFound)
	}
}
