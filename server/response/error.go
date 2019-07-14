package response

import (
	"database/sql"
	"log"
	"net/http"
)

// create a switch statement to return status code and error message

// SendError writes a struct to the writer
func SendError(w http.ResponseWriter, err error, statusCode int) {
	log.Print("Send Error", err)
	if err == sql.ErrNoRows {
		http.Error(w, "No rows found", statusCode)
		return
	}
	http.Error(w, err.Error(), statusCode)
	return
}
