package server

import (
	"context"
	"encoding/json"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//TODO
//write handlers
//note that these will be JSON to begin but eventually I want to use html templates

// handler for the index
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["msg"] = "Welcome to Flex Creek!"

	writeJSON(w, http.StatusOK, msg)
}

// handler for wod creation
func (s *Server) handleCreateWod(w http.ResponseWriter, r *http.Request) {
	//making a generic context for now, although this could be something different later

	ctx := context.Background()

	var wod *flexcreek.Wod

	err := json.NewDecoder(r.Body).Decode(&wod)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
		return
	}

	id := uuid.New().String()

	wod.ID = id

	err = s.WodService.CreateWod(ctx, wod)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, wod)

}

func (s *Server) handleGetAllWods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	wods, err := s.WodService.GetAllWods(ctx)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, wods)
}

func (s *Server) handleGetRandomWod(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	wod, err := s.WodService.GetRandomWod(ctx)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, wod)
}

func (s *Server) handleGetWodbyID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	id := vars["wodID"]

	wod, err := s.WodService.GetWodByID(ctx, id)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err)
	}

	writeJSON(w, http.StatusOK, wod)
}

// helper funcs
func writeJSON(w http.ResponseWriter, statusCode int, v any) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(v)

}
