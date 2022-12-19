package server

import (
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

//add Run method
