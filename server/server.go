package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/gorilla/mux"
)

type Server struct {
	Router      *mux.Router
	Srvr        *http.Server
	WodService  flexcreek.WodService
	UserService flexcreek.UserService
}

func NewServer(r *mux.Router, ws flexcreek.WodService, us flexcreek.UserService) *Server {

	listenAddr := ":8080"

	return &Server{
		Router: r,
		Srvr: &http.Server{
			Addr: listenAddr,
		},
		WodService:  ws,
		UserService: us,
	}
}

// function to handle the index route
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["msg"] = "Welcome to Flex Creek!"

	writeJSON(w, http.StatusOK, msg)
}

// add Register Routes method
func (s *Server) registerRoutes() {
	//index
	s.Router.HandleFunc("/", handleIndex).Methods("GET")

	//wod handlers
	s.Router.HandleFunc("/wod", s.handleCreateWod).Methods("POST")
	s.Router.HandleFunc("/wod", s.handleGetAllWods).Methods("GET")
	s.Router.HandleFunc("/randomwod", s.handleGetRandomWod).Methods("GET")
	s.Router.HandleFunc("/wod/{wodID}", s.handleGetWodbyID).Methods("GET")
	s.Router.HandleFunc("/wod/type/{wodType}", s.handleGetWodbyType).Methods("GET")
	s.Router.HandleFunc("/wod/update/{wodID}", s.handleUpdateWod).Methods("POST")
	s.Router.HandleFunc("/wod/delete/{wodID}", s.handleDeleteWod).Methods("DELETE")

	//user handlers
	s.Router.HandleFunc("/user", s.handleCreateUser).Methods("POST")
}

// add Run method
func (s *Server) Run() {

	s.registerRoutes()

	fmt.Printf("Flex Creek running on port %s", s.Srvr.Addr)

	s.Srvr.Handler = s.Router

	s.Srvr.ListenAndServe()
}

// helpers
func writeJSON(w http.ResponseWriter, statusCode int, v any) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(v)

}
