package app

import (
	"encoding/json"
	// post gres
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

// Message data object
type Message struct {
	ID          string     `json:"id"`
	MessageText string     `json:"message_text"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Archived    bool       `json:"archived"`
	UserID      int        `json:"user_id"`
	GameID      *int       `json:"game_id"`
}

// MessageInput for incoming inserted messages
type MessageInput struct {
	MessageText string `json:"message_text"`
	Archived    bool   `json:"archived"`
	UserID      string `json:"user_id"`
	GameID      string `json:"game_id"`
}

// Messages is a slice of message.
type Messages struct {
	Messages []Message
}

// GetMessages gets all the messages. This should probably be on a per...game basis
func GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := Messages{}

	err := queryMessages(&messages)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(messages)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}

func queryMessages(messages *Messages) error {
	rows, err := DB.Query(`SELECT * FROM messages;`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		message := Message{}
		err = rows.Scan(
			&message.ID,
			&message.MessageText,
			&message.CreatedAt,
			&message.UpdatedAt,
			&message.Archived,
			&message.UserID,
			&message.GameID,
		)
		if err != nil {
			return err
		}
		messages.Messages = append(messages.Messages, message)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// InsertMessage inserts a message to the database
func InsertMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var messageInput MessageInput
	err := decoder.Decode(&messageInput)
	if err != nil {
		log.Print(err)
	}
	log.Println(messageInput)
}

func insertMessageQuery() {
	var err error
	insertStatement := `
	INSERT INTO messages (message_text, created_at, user_id)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err = DB.QueryRow(insertStatement, "Second Message", time.Now(), 1).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	log.Println("New record ID is:", id)

}
