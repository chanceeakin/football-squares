package controller

import (
	"log"
	"net/http"
)

// FourOhFourHandler is the catch all
func FourOhFourHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("404")
	http.Error(w, `Y u do this`, http.StatusNotFound)
}
