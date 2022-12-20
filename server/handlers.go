package server

import (
	"context"
	"encoding/json"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
)

//TODO
//write handlers
//note that these will be JSON to begin but eventually I want to use html templates

// handler for the index
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string)

	msg["msg"] = "Welcome to Flex Creek!"

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(msg)
}

// handler for wod creation
func (s *Server) handleCreateWod(w http.ResponseWriter, r *http.Request) {
	//making a generic context for now, although this could be something different later

	ctx := context.Background()

	var wod *flexcreek.Wod

	err := json.NewDecoder(r.Body).Decode(&wod)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	id := uuid.New().String()

	wod.ID = id

	err = s.WodService.CreateWod(ctx, wod)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(wod)

}

func (s *Server) handleGetAllWods(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	wods, err := s.WodService.GetAllWods(ctx)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(wods)
}

//todo
//write helper function to return json response
//write helper function to retrun json response if err
