package server

import (
	"context"
	"encoding/json"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/ekholme/flex_creek/middleware"
	"github.com/ekholme/flex_creek/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user *flexcreek.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	id := uuid.New().String()

	user.ID = id

	err = s.UserService.CreateUser(ctx, user)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	users, err := s.UserService.GetAllUsers(ctx)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := mux.Vars(r)["userID"]

	user, err := s.UserService.GetUserByID(ctx, id)

	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

// TODO
func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var l *flexcreek.Login

	err := json.NewDecoder(r.Body).Decode(&l)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()

	ref, err := s.UserService.GetUserByUsername(ctx, l.Username)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	//ensure pw match
	if err = validateLogin(l, ref); err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, err)
		return
	}

	//next, create auth from user info, generate a token, create a cookie, and set cookie
	a := middleware.CreateAuth(ref)

	if middleware.GenerateToken(a); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	cookie := http.Cookie{
		Name:     "FLEXAUTH",
		Value:    a.Token,
		Path:     "/",
		MaxAge:   7200,
		HttpOnly: true,
		Secure:   false, //for now while testing
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	utils.WriteJSON(w, http.StatusOK, "logged in!")
}

// welcome handler to ensure auth is working
func (s *Server) handleWelcome(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	msg := "Welcome " + claims.Username

	utils.WriteJSON(w, http.StatusOK, msg)
}

// helper to ensure passwords match
func validateLogin(l *flexcreek.Login, u *flexcreek.User) error {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))

	if err != nil {
		return err
	}

	return nil
}
