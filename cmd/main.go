package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const listenAddr string = ":8080"

// just a little hello world
func main() {

	//need to add defer client.Close() somewhere for firestore client
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	s := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}

	fmt.Println("Welcome to Flex Creek!")

	s.ListenAndServe()
}

// handler for the index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	hw := make(map[string]string)

	hw["hello"] = "world"

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(hw)
}
