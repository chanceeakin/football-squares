package controller

import (
	"encoding/json"
	user "github.com/chanceeakin/football-squares/server/models/user"
	"log"
	"net/http"
)

// GetUsers gets all the messages. This should probably be on a per...game basis
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userMap := user.Users{}

	err := user.QueryUsers(&userMap)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(userMap)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}
