package main

import (
	"log"

	app "github.com/chanceeakin/football-squares/server/app"
)

func main() {
	log.Println("Starting service")
	app.Run()
}
