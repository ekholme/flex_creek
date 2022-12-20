package server

import (
	"fmt"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/gorilla/mux"
)

//TODO
//implement server logic

type Server struct {
	Router      *mux.Router
	Srvr        *http.Server
	WodService  flexcreek.WodService
	UserService flexcreek.UserService
}

func NewServer(r *mux.Router, ws flexcreek.WodService, us flexcreek.UserService) *Server {
	return &Server{
		Router:      r,
		Srvr:        &http.Server{},
		WodService:  ws,
		UserService: us,
	}
}

// add Register Routes method
func (s *Server) registerRoutes() {
	s.Router.HandleFunc("/", handleIndex).Methods("GET")
	s.Router.HandleFunc("/wod", s.handleCreateWod).Methods("POST")
	s.Router.HandleFunc("/wod", s.handleGetAllWods).Methods("GET")
}

// add Run method
func (s *Server) Run() {
	listenAddr := ":8080"

	s.registerRoutes()

	fmt.Printf("Flex Creek running on port %s", listenAddr)

	s.Srvr.Addr = listenAddr
	s.Srvr.Handler = s.Router

	s.Srvr.ListenAndServe()
}
