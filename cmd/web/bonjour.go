package web

import (
	"fmt"
	"net/http"
)

func BonjourWebHandler(w http.ResponseWriter, r *http.Request) {
	message := "Bonjour"

	// Write the message to the response
	_, err := fmt.Fprintln(w, message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
