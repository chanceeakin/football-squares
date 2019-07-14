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

// var routes = Routes{
// 	Route{
// 		"Messages",
// 		"/messages",
// 		controllers.GetMessages,
// 	},
// 	Route{
// 		"Message",
// 		"/message",
// 		controllers.MessageHandler,
// 	},
// 	Route{
// 		"Games",
// 		"/games",
// 		controllers.GetGames,
// 	},
// 	Route{
// 		"Game",
// 		"/game",
// 		controllers.GameHandler,
// 	},
// 	Route{
// 		"Messages By Game",
// 		"/games/messages",
// 		controllers.MessageByGameHandler,
// 	},
// 	Route{
// 		"Users",
// 		"/users",
// 		controllers.UsersHandlers,
// 	},
// 	Route{
// 		"User",
// 		"/user",
// 		controllers.UserHandlers,
// 	},
// 	Route{
// 		"Index",
// 		"/",
// 		controllers.ErrorHandler,
// 	},
// }
