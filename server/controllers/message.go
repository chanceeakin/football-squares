package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	message "football-squares/server/models/message"
	response "football-squares/server/response"
	// post gres
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// GetMessages gets all the messages. This should probably be on a per...game basis
func GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `Not Found`, http.StatusNotFound)
		return
	}
	messagesArr := message.Messages{}

	err := message.QueryMessages(&messagesArr)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, messagesArr)
}

// PostMessage inserts a message to the database
func postMessage(w http.ResponseWriter, r *http.Request) {
	var err error
	var messageInput message.Input
	var out common.ID
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&messageInput)

	if err != nil {
		log.Print(err)
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	out, err = message.PostMessageQuery(&messageInput)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}
	response.SendJSON(w, out)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	var input common.ID
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	val, err1 := message.QueryMessage(&input)
	if err1 != nil {
		log.Print(err1)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, val)
}

// MessageHandler is the switch for REST Methods
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessage(w, r)
	case http.MethodPost:
		postMessage(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, `Not Found`, http.StatusNotFound)
	}
}
