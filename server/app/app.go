package app

import (
	"log"
	"net/http"

	controllers "football-squares/server/controllers"
	db "football-squares/server/db"
	middleware "football-squares/server/middleware"
	"github.com/gorilla/mux"
)

// Run the app
func Run(d *db.InitData) {
	uriString := d.Host + ":" + d.Port
	db.Init(d)
	defer db.CleanUp()

r := mux.NewRouter()
	r.Handle("/messages", middleware.Logger(http.HandlerFunc(controllers.GetMessages)))
	r.Handle("/message", middleware.Logger(http.HandlerFunc(controllers.MessageHandler)))
	r.Handle("/games", middleware.Logger(http.HandlerFunc(controllers.GetGames)))
	r.Handle("/game", middleware.Logger(http.HandlerFunc(controllers.GameHandler)))
	r.Handle("/game/messages", middleware.Logger(http.HandlerFunc(controllers.MessageByGameHandler)))
	r.Handle("/users", middleware.Logger(http.HandlerFunc(controllers.UsersHandlers)))
	r.Handle("/user", middleware.Logger(http.HandlerFunc(controllers.UserHandlers)))
	r.Handle("/", middleware.Logger(http.HandlerFunc(controllers.ErrorHandler)))
	log.Fatal(http.ListenAndServe(uriString, r))

}
