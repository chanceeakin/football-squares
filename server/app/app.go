package app

import (
	"log"
	"net/http"

	controllers "github.com/chanceeakin/football-squares/server/controllers"
	db "github.com/chanceeakin/football-squares/server/db"
	middleware "github.com/chanceeakin/football-squares/server/middleware"
)

// Run the app
func Run(d *db.InitData) {
	uriString := d.Host + ":" + d.Port
	db.Init(d)
	defer db.CleanUp()
	http.Handle("/messages", middleware.Logger(http.HandlerFunc(controllers.GetMessages)))
	http.Handle("/games", middleware.Logger(http.HandlerFunc(controllers.GetGames)))
	http.Handle("/game", middleware.Logger(http.HandlerFunc(controllers.GetGame)))
	http.Handle("/users", middleware.Logger(http.HandlerFunc(controllers.GetUsers)))
	http.Handle("/new-message", middleware.Logger(http.HandlerFunc(controllers.PostMessage)))
	log.Fatal(http.ListenAndServe(uriString, nil))

}
