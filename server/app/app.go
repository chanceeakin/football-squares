package app

import (
	"log"
	"net/http"

	controllers "football-squares/server/controllers"
	db "football-squares/server/db"
	routes "football-squares/server/routes"
	"github.com/urfave/negroni"
)

// UGH YES.
// https://thenewstack.io/make-a-restful-json-api-go/

// Routes is the slice of those configs
var Routes = make([]routes.Route, 0)

func init() {
	Routes = append(Routes, controllers.GameRoutes()...)
	Routes = append(Routes, controllers.MessageRoutes()...)
	Routes = append(Routes, controllers.UserRoutes()...)
}

// Run the app
func Run(d *db.InitData) {
	uriString := d.Host + ":" + d.Port
	db.Init(d)
	defer db.CleanUp()

	r := routes.NewRouter(Routes)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(uriString, n))

}
