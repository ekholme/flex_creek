package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/ekholme/flex_creek/middleware"
	"github.com/ekholme/flex_creek/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// handler for wod creation
func (s *Server) handleCreateWod(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	claims := r.Context().Value("flexclaims").(*middleware.CustomClaims)

	var wod *flexcreek.Wod

	err := json.NewDecoder(r.Body).Decode(&wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	id := uuid.New().String()

	wod.ID = id

	wod.AddedBy = claims.Username

	validate := validator.New()

	err = validate.Struct(wod)

	//I should probably clean this up later to give more specific errors
	//but this is ok for now i think
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.WodService.CreateWod(ctx, wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wod)

}

func (s *Server) handleGetAllWods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	wods, err := s.WodService.GetAllWods(ctx)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wods)
}

func (s *Server) handleGetRandomWod(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	wod, err := s.WodService.GetRandomWod(ctx)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wod)
}

func (s *Server) handleGetWodbyID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	id := vars["wodID"]

	wod, err := s.WodService.GetWodByID(ctx, id)

	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wod)
}

func (s *Server) handleGetWodbyType(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	t := vars["wodType"]

	wods, err := s.WodService.GetWodsbyType(ctx, t)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wods)
}

func (s *Server) handleUpdateWod(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	id := vars["wodID"]

	var wod *flexcreek.Wod

	err := json.NewDecoder(r.Body).Decode(&wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uWod, err := s.WodService.UpdateWod(ctx, id, wod)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, uWod)

}

func (s *Server) handleDeleteWod(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	id := vars["wodID"]

	err := s.WodService.DeleteWod(ctx, id)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	msg := make(map[string]string)

	msg["message"] = "Wod Deleted"

	utils.WriteJSON(w, http.StatusOK, msg)

}

// wip -- getting wods by query
func (s *Server) handleGetWodsByQuery(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	urlParams := r.URL.Query()

	m := make(map[string]string, len(urlParams))

	for i, v := range urlParams {
		if len(v) > 1 {
			utils.WriteJSON(w, http.StatusBadRequest, errors.New("cannot handle duplicate parameters"))
			return
		}

		s := strings.Join(v, "")

		m[i] = s
	}

	wods, err := s.WodService.GetWodsByQuery(ctx, m)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, wods)
}
