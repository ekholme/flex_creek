package server

import (
	"encoding/json"
	"net/http"
)

//TODO
//write handlers
//note that these will be JSON to begin but eventually I want to use html templates

// handler for the index
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["msg"] = "Welcome to Flex Creek!"

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(msg)
}

// handler for wod creation
func handleCreateWod(w http.ResponseWriter, r *http.Request) {

}
