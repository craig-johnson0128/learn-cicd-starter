package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON
	dat, err := json.Marshal(payload)
	if err != nil {
		// Log the error and send a 500 Internal Server Error
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")
	// Write the status code to the response
	w.WriteHeader(code)
	// Write the JSON data to the response
	if _, err := w.Write(dat); err != nil {
		// Handle errors that occur while writing the response
		log.Printf("Error writing response: %s", err)
	}
}
