package server

import (
	"context"
	"net/http"

	"github.com/ekholme/flex_creek/middleware"
	"github.com/ekholme/flex_creek/utils"
	"github.com/gorilla/mux"
)

// handlers for favorite methods
func (s *Server) handleCreateFavorite(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	ctx := context.Background()

	vars := mux.Vars(r)

	wid := vars["wodID"]

	wod, err := s.WodService.GetWodByID(ctx, wid)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.FavoriteService.CreateFavoriteWod(ctx, claims.UserID, wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	msg := wod.Name + " added to favorites"

	utils.WriteJSON(w, http.StatusOK, msg)
}

func (s *Server) handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	ctx := context.Background()

	vars := mux.Vars(r)

	wid := vars["wodID"]

	err := s.FavoriteService.DeleteFavoriteWod(ctx, claims.UserID, wid)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	//eventually might want to include the wod name in this msg to the user
	msg := "wod removed from favorites"

	utils.WriteJSON(w, http.StatusOK, msg)
}

func (s *Server) handleGetAllFavorites(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	ctx := context.Background()

	wods, err := s.FavoriteService.GetAllFavoriteWods(ctx, claims.UserID)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, wods)

}
