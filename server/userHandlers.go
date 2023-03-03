package server

import (
	"context"
	"encoding/json"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
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

//RESUME HERE
