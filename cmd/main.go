package main

import (
	"github.com/ekholme/flex_creek/frstr"
	"github.com/ekholme/flex_creek/server"
	"github.com/gorilla/mux"
)

func main() {

	client := frstr.NewClient()

	defer client.Close()

	ws := frstr.NewWodService(client)
	us := frstr.NewUserService(client)
	r := mux.NewRouter()

	s := server.NewServer(r, ws, us)

	s.Run()
}
