package server

import "net/http"

// handlers for favorite methods
func (s *Server) handleCreateFavorite(w http.ResponseWriter, r *http.Request) {
	//RESUME HERE
	//i think i need to implement auth before i can do this, bc i'll want
	//to grab the userid from a jwt and then pass that into the createFavorite
	//method in the favoriteService
	//then from there I can decode a json body with the wod
}

func (s *Server) handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleGetAllFavorites(w http.ResponseWriter, r *http.Request) {

}
