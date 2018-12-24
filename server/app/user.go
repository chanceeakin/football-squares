package app

import (
	"encoding/json"
	"log"
	"net/http"
)

// User is a data struct for a given user
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Users is a slice of user.
type Users struct {
	Users []User
}

// GetUsers gets all the messages. This should probably be on a per...game basis
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{}

	err := queryUsers(&users)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
		http.Error(w, `Internal Error`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	w.Write(out)
}

func queryUsers(users *Users) error {
	rows, err := DB.Query(`SELECT * FROM users;`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		if err != nil {
			return err
		}
		users.Users = append(users.Users, user)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
