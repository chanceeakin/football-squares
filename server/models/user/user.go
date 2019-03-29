package user

import (
	db "github.com/chanceeakin/football-squares/server/db"
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

// QueryUsers queries the db for existing users.
func QueryUsers(users *Users) error {
	rows, err := db.DB.Query(`SELECT * FROM users;`)
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
