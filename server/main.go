package main

import (
	"log"

	app "github.com/chanceeakin/football-squares/server/app"
	db "github.com/chanceeakin/football-squares/server/db"
)

const (
	host     = "localhost"
	port     = "8000"
	dbPort   = 5432
	user     = "postgres"
	password = "password"
	dbname   = "football_square"
)

func main() {

	log.Println("Starting service")
	initData := db.InitData{Host: host, Port: port, User: user, Password: password, Dbname: dbname, DbPort: dbPort}

	app.Run(&initData)
}
