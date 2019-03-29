package controller

import (
	"encoding/json"
	message "github.com/chanceeakin/football-squares/server/models/message"
	// post gres
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// GetMessages gets all the messages. This should probably be on a per...game basis
func GetMessages(w http.ResponseWriter, r *http.Request) {
	messagesArr := message.Messages{}

	err := message.QueryMessages(&messagesArr)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(messagesArr)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}

// PostMessage inserts a message to the database
func PostMessage(w http.ResponseWriter, r *http.Request) {
	var err error
	var messageInput message.Input
	var out message.Out
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&messageInput)

	if err != nil {
		log.Print(err)
	}

	out, err = message.PostMessageQuery(&messageInput)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}
	out1, err := json.Marshal(out)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out1)
}
