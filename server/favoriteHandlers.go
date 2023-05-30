package server

import (
	"context"
	"net/http"

	"github.com/ekholme/flex_creek/middleware"
	"github.com/gorilla/mux"
)

// handlers for favorite methods
func (s *Server) handleCreateFavorite(w http.ResponseWriter, r *http.Request) {
	//RESUME HERE
	//i think i need to implement auth before i can do this, bc i'll want
	//to grab the userid from a jwt and then pass that into the createFavorite
	//method in the favoriteService
	//then from there I can decode a json body with the wod

	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	ctx := context.Background()

	vars := mux.Vars(r)

	wid := vars["wodID"]

	wod, err := s.WodService.GetWodByID(ctx, wid)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}

	err = s.FavoriteService.CreateFavoriteWod(ctx, claims.UserID, wod)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}

	msg := wod.Name + " added to favorites"

	writeJSON(w, http.StatusOK, msg)
}

func (s *Server) handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleGetAllFavorites(w http.ResponseWriter, r *http.Request) {

}
