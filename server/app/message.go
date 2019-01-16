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
	UserID      string     `json:"user_id"`
	GameID      *string    `json:"game_id"`
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

// MessageOut is the message output for an insert
type MessageOut struct {
	ID string
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
	var err error
	var messageInput MessageInput
	err = decoder.Decode(&messageInput)
	out := MessageOut{}

	if err != nil {
		log.Print(err)
	}

	out, err = insertMessageQuery(&messageInput)
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

func insertMessageQuery(messageInput *MessageInput) (MessageOut, error) {
	var err error
	insertStatement := `
	INSERT INTO messages (message_text, created_at, user_id, game_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	out := MessageOut{}
	err = DB.QueryRow(insertStatement, &messageInput.MessageText, time.Now(), &messageInput.UserID, &messageInput.GameID).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}
