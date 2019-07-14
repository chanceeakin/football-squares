package main

import (
	"log"

	app "football-squares/server/app"
	db "football-squares/server/db"
)

const (
	host     = "0.0.0.0"
	port     = "8080"
	dbHost   = "postgres"
	dbPort   = 5432
	user     = "postgres"
	password = "password"
	dbname   = "football_square"
)

func main() {

	log.Println("Starting football services app")
	initData := db.InitData{Host: host, Port: port, User: user, Password: password, Dbname: dbname, DbHost: dbHost, DbPort: dbPort}

	app.Run(&initData)
}
