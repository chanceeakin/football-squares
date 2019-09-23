package message

import (
	common "football-squares/server/common"
	db "football-squares/server/db"
	"log"
	"time"
)

const insertOneSQL = `INSERT INTO messages (message_text, created_at, user_id, game_id)
	VALUES ($1, $2, $3, $4)
  RETURNING id;`
const selectFromGameSQL = `SELECT * FROM messages where game_id=$1;`
const selectAllSQL = `SELECT * FROM messages;`

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
	MessageText string `json:"message_text" validate:"required"`
	Archived    bool   `json:"archived" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
	GameID      string `json:"game_id" validate:"required"`
}

// Messages is a slice of message.
type Messages struct {
	Messages []Message
}

// PostMessageQuery posts a single message
func PostMessageQuery(messageInput *Input) (common.ID, error) {
	var err error
	out := common.ID{}
	err = db.DB.QueryRow(insertOneSQL, &messageInput.MessageText, time.Now(), &messageInput.UserID, &messageInput.GameID).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}

// QueryMessages returns all messages
func QueryMessages(messages *Messages) error {
	rows, err := db.DB.Query(selectAllSQL)
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

// QueryMessagesByGame returns a games' worth of messages
func QueryMessagesByGame(i *common.ID) (*Messages, error) {
	val := Messages{}
	rows, err := db.DB.Query(selectFromGameSQL, &i.ID)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		val.Messages = append(val.Messages, message)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// QueryMessage returns a single message
func QueryMessage(i *common.ID) (Message, error) {
	val := Message{}
	row := db.DB.QueryRow(`SELECT * FROM messages where id=$1;`, &i.ID)
	err := row.Scan(
		&val.ID,
		&val.MessageText,
		&val.CreatedAt,
		&val.UpdatedAt,
		&val.Archived,
		&val.UserID,
		&val.GameID,
	)
	if err != nil {
		log.Print(err)
		return val, err
	}
	return val, nil
}
