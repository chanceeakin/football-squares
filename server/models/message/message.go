package message

import (
	db "github.com/chanceeakin/football-squares/server/db"
	"log"
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

// Input for incoming inserted messages
type Input struct {
	MessageText string `json:"message_text"`
	Archived    bool   `json:"archived"`
	UserID      string `json:"user_id"`
	GameID      string `json:"game_id"`
}

// Messages is a slice of message.
type Messages struct {
	Messages []Message
}

// Out is the message output for an insert
type Out struct {
	ID string
}

func PostMessageQuery(messageInput *Input) (Out, error) {
	var err error
	insertStatement := `
	INSERT INTO messages (message_text, created_at, user_id, game_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	out := Out{}
	err = db.DB.QueryRow(insertStatement, &messageInput.MessageText, time.Now(), &messageInput.UserID, &messageInput.GameID).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}

// QueryMessages returns all messages
func QueryMessages(messages *Messages) error {
	rows, err := db.DB.Query(`SELECT * FROM messages;`)
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
