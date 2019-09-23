package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is the config struct for api routing
type Route struct {
	Name        string
	Path        string
	HandlerFunc http.HandlerFunc
	Method      string
}

// NewRouter Returns new router
func NewRouter(routes []Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		router.
			Path(route.Path).
			Name(route.Name).
			Methods(route.Method).
			Handler(route.HandlerFunc)
	}

	return router
}
