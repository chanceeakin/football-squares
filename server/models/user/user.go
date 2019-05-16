package user

import (
	db "football-squares/server/db"
	"log"
	"golang.org/x/crypto/bcrypt"
)

// User is a data struct for a given user
type User struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	IsAdmin   *bool   `json:"is_admin"`
}

// NoPasswordUser scrubs out the password from the return value
type NoPasswordUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsAdmin   *bool  `json:"is_admin"`
}

// Users is a slice of user.
type Users struct {
	Users []NoPasswordUser
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
	Password string `json:"password"`
}

// QueryUsers queries the db for existing users.
func QueryUsers(users *Users) error {
	rows, err := db.DB.Query(`SELECT id, first_name, last_name, email, is_admin FROM users;`)
	if err != nil {
		log.Print(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		user := NoPasswordUser{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.IsAdmin,
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
func QueryUser(input *GetInput) (NoPasswordUser, error) {
	returnUser := NoPasswordUser{}
	row := db.DB.QueryRow(`SELECT id, first_name, last_name, email, is_admin FROM users where id=$1;`, &input.ID)
	err := row.Scan(
		&returnUser.ID,
		&returnUser.FirstName,
		&returnUser.LastName,
		&returnUser.Email,
		&returnUser.IsAdmin,
	)
	if err != nil {
		log.Print(err)
		return returnUser, err
	}
	return returnUser, nil
}

func hashAndSalt(pwd string) string {

    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost.
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

// InsertUser inserts a user into the DB
func InsertUser(input *Input) (Out, error) {
	var err error
	hashedPassword := hashAndSalt(input.Password)
	insertStatement := `
	INSERT INTO users (first_name, last_name, email, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	out := Out{}
	err = db.DB.QueryRow(insertStatement, &input.FirstName, &input.LastName, &input.Email, &hashedPassword).Scan(&out.ID)
	if err != nil {
		log.Print(err)
		return out, err
	}
	log.Println("New record ID is:", out.ID)
	return out, nil
}
