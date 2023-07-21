package server

import (
	"fmt"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/ekholme/flex_creek/middleware"
	"github.com/ekholme/flex_creek/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	Router          *mux.Router
	Srvr            *http.Server
	WodService      flexcreek.WodService
	UserService     flexcreek.UserService
	FavoriteService flexcreek.FavoriteService
}

func NewServer(r *mux.Router, ws flexcreek.WodService, us flexcreek.UserService, fs flexcreek.FavoriteService) *Server {

	listenAddr := ":8080"

	return &Server{
		Router: r,
		Srvr: &http.Server{
			Addr: listenAddr,
		},
		WodService:      ws,
		UserService:     us,
		FavoriteService: fs,
	}
}

// function to handle the index route
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["msg"] = "Welcome to Flex Creek!"

	utils.WriteJSON(w, http.StatusOK, msg)
}

// add Register Routes method
func (s *Server) registerRoutes() {
	//no auth routes
	s.Router.HandleFunc("/", handleIndex).Methods("GET")
	s.Router.HandleFunc("/login", s.handleLogin).Methods("POST")
	s.Router.HandleFunc("/register", s.handleCreateUser).Methods("POST")

	authRouter := s.Router.PathPrefix("/o").Subrouter()
	authRouter.Use(middleware.JWTMiddlewareTwo)

	//wod handlers
	authRouter.HandleFunc("/wod", s.handleCreateWod).Methods("POST")
	authRouter.HandleFunc("/wod", s.handleGetAllWods).Methods("GET")
	authRouter.HandleFunc("/randomwod", s.handleGetRandomWod).Methods("GET")
	//these routes are janky -- look into how people usually do this
	authRouter.HandleFunc("/wod/{wodID}", s.handleGetWodbyID).Methods("GET")
	authRouter.HandleFunc("/wod/type/{wodType}", s.handleGetWodbyType).Methods("GET")
	authRouter.HandleFunc("/wodquery", s.handleGetWodsByQuery).Methods("GET")
	authRouter.HandleFunc("/wod/update/{wodID}", s.handleUpdateWod).Methods("POST")
	authRouter.HandleFunc("/wod/delete/{wodID}", s.handleDeleteWod).Methods("DELETE")

	//user handlers
	//most of the user management needs to be for admins
	authRouter.HandleFunc("/user/{userID}", s.handleGetUserByID).Methods("GET")
	authRouter.HandleFunc("/users", s.handleGetAllUsers).Methods("GET")
	authRouter.HandleFunc("/user/{userID}/delete", s.handleDeleteUser).Methods("GET")
	//welcome
	authRouter.HandleFunc("/o/welcome", middleware.JWTMiddleware(s.handleWelcome)).Methods("GET")

	//favorites
	authRouter.HandleFunc("/wod/{wodID}/favorite", middleware.JWTMiddleware(s.handleCreateFavorite)).Methods("POST")
	authRouter.HandleFunc("/favoritewods", middleware.JWTMiddleware(s.handleGetAllFavorites)).Methods("GET")
	authRouter.HandleFunc("/wod/{wodID}/deletefavorite", middleware.JWTMiddleware(s.handleDeleteFavorite)).Methods("GET")
}

// add Run method
func (s *Server) Run() {

	s.registerRoutes()

	fmt.Printf("Flex Creek running on port %s", s.Srvr.Addr)

	s.Srvr.Handler = s.Router

	s.Srvr.ListenAndServe()
}
