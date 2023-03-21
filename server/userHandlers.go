package server

import (
	"context"
	"encoding/json"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user *flexcreek.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}

	id := uuid.New().String()

	user.ID = id

	err = s.UserService.CreateUser(ctx, user)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *Server) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	users, err := s.UserService.GetAllUsers(ctx)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, users)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := mux.Vars(r)["userID"]

	user, err := s.UserService.GetUserByID(ctx, id)

	if err != nil {
		writeJSON(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// TODO
func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {

}
