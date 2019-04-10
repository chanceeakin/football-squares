package db

import (
	"database/sql"
	"fmt"
	"log"
	// postgres connection
	_ "github.com/lib/pq"
)

// InitData is the initial data needed for a connection to teh db
type InitData struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	DbHost   string
	DbPort   int
}

//DB stores the databse connection in memory
var DB *sql.DB

// Init initializes the postgres conneciton.
func Init(d *InitData) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.DbHost, d.DbPort, d.User, d.Password, d.Dbname)
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

// CleanUp cleans up the database
func CleanUp() {
	log.Println("Cleaning up database connections")
	DB.Close()
}
