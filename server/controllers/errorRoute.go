package controller

import (
	"log"
	"net/http"
)

// ErrorHandler is the catch all
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("404")
	http.Error(w, `Y u do this`, http.StatusNotFound)
}
