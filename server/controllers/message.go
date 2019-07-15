package controller

import (
	"encoding/json"
	common "football-squares/server/common"
	message "football-squares/server/models/message"
	response "football-squares/server/response"
	routes "football-squares/server/routes"
	// post gres
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// MessageRoutes is the declaration for all routes
func MessageRoutes() []routes.Route {
	messageRoutes := make([]routes.Route, 4)
	messageRoutes = append(messageRoutes, routes.Route{
		Name:        "Messages",
		Path:        "/messages",
		HandlerFunc: getMessages,
		Method:      "GET",
	},
		routes.Route{
			Name:        "Message",
			Path:        "/message",
			HandlerFunc: getMessage,
			Method:      "GET",
		},
		routes.Route{
			Name:        "Post Message",
			Path:        "/message",
			HandlerFunc: postMessage,
			Method:      "POST",
		},
		routes.Route{
			Name:        "Messages by Game",
			Path:        "/messages/game",
			HandlerFunc: messageByGameHandler,
			Method:      "GET",
		})
	return messageRoutes
}

func getMessages(w http.ResponseWriter, r *http.Request) {
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

//MessageByGameHandler handles returning messages per game
func messageByGameHandler(w http.ResponseWriter, r *http.Request) {
	var input common.ID
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		log.Print("Message By Game not post")
		http.Error(w, `Wrong Method`, http.StatusMethodNotAllowed)
	}
	messageArr, err := message.QueryMessagesByGame(&input)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, messageArr)

}
