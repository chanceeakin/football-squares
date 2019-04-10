package user

import (
	db "football-squares/server/db"
	"log"
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

// GetInput is the input type for finding a single user.
type GetInput struct {
	ID string `json:"id"`
}

// Out is the message output for an insert
type Out struct {
	ID string
}

//Input is for inserting a user
type Input struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// QueryUsers queries the db for existing users.
func QueryUsers(users *Users) error {
	rows, err := db.DB.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		return err
	}
	return nil
}

// QueryUser finds a single user
func QueryUser(input *GetInput) (User, error) {
	returnUser := User{}
	row := db.DB.QueryRow(`SELECT * FROM users where id=$1;`, &input.ID)
	err := row.Scan(
		&returnUser.ID,
		&returnUser.FirstName,
		&returnUser.LastName,
		&returnUser.Email,
	)
	if err != nil {
		log.Print(err)
		return returnUser, err
	}
	return returnUser, nil
}

// InsertUser inserts a user into the DB
func InsertUser(input *Input) (Out, error) {
	var err error
	insertStatement := `
	INSERT INTO users (first_name, last_name, email)
	VALUES ($1, $2, $3)
	RETURNING id`
	out := Out{}
	err = db.DB.QueryRow(insertStatement, &input.FirstName, &input.LastName, &input.Email).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}
