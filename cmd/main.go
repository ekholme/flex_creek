package main

import (
	"github.com/ekholme/flex_creek/frstr"
	"github.com/ekholme/flex_creek/server"
	"github.com/gorilla/mux"
)

// just a little hello world
func main() {
	//add defer client.Close here somewhere

	client := frstr.NewClient()
	ws := frstr.NewWodService(client)
	r := mux.NewRouter()

	s := server.NewServer(r, ws, nil)

	s.Run()
}
