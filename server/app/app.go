package app

import (
	"log"
	"net/http"

	controllers "football-squares/server/controllers"
	db "football-squares/server/db"
	middleware "football-squares/server/middleware"
)

// Run the app
func Run(d *db.InitData) {
	uriString := d.Host + ":" + d.Port
	db.Init(d)
	defer db.CleanUp()
	http.Handle("/messages", middleware.Logger(http.HandlerFunc(controllers.GetMessages)))
	http.Handle("/message", middleware.Logger(http.HandlerFunc(controllers.MessageHandler)))
	http.Handle("/games", middleware.Logger(http.HandlerFunc(controllers.GetGames)))
	http.Handle("/game", middleware.Logger(http.HandlerFunc(controllers.GameHandler)))
	http.Handle("/game/messages", middleware.Logger(http.HandlerFunc(controllers.MessageByGameHandler)))
	http.Handle("/users", middleware.Logger(http.HandlerFunc(controllers.UsersHandlers)))
	http.Handle("/user", middleware.Logger(http.HandlerFunc(controllers.UserHandlers)))
	http.Handle("/", middleware.Logger(http.HandlerFunc(controllers.ErrorHandler)))
	log.Fatal(http.ListenAndServe(uriString, nil))

}
