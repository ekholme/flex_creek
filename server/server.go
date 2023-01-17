package server

import (
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

// add Register Routes method
func (s *Server) registerRoutes() {
	s.Router.HandleFunc("/", handleIndex).Methods("GET")
	s.Router.HandleFunc("/wod", s.handleCreateWod).Methods("POST")
	s.Router.HandleFunc("/wod", s.handleGetAllWods).Methods("GET")
	s.Router.HandleFunc("/randomwod", s.handleGetRandomWod).Methods("GET")
	s.Router.HandleFunc("/wod/{wodID}", s.handleGetWodbyID).Methods("GET")
}

// add Run method
func (s *Server) Run() {

	s.registerRoutes()

	fmt.Printf("Flex Creek running on port %s", s.Srvr.Addr)

	s.Srvr.Handler = s.Router

	s.Srvr.ListenAndServe()
}
