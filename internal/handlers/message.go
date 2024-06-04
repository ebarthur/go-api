package handlers

import (
	"encoding/json"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Welcome to the API"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Prepare the data to be encoded into JSON
	data := map[string]string{
		"message": "Welcome to the API",
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the data into JSON and send it as a response
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
