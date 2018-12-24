package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// postgres connection
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "football_square"
)

//DB stores the databse connection in memory
var DB *sql.DB

// DBInit initializes the postgres conneciton.
func DBInit() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Print(err)
	}
	log.Println("Successfully connected!")
}

// DBCleanUp cleans up the database
func DBCleanUp() {
	log.Println("Cleaning up database connections")
	DB.Close()
}

// Run the app
func Run() {
	DBInit()
	defer DBCleanUp()
	http.HandleFunc("/messages", GetMessages)
	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/new-message", InsertMessage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
