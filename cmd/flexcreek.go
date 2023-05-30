package main

import (
	"log"

	"github.com/ekholme/flex_creek/frstr"
	"github.com/ekholme/flex_creek/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Couldn't load .env file")
	}

	client := frstr.NewFirestoreClient()

	defer client.Close()

	ws := frstr.NewWodService(client)
	us := frstr.NewUserService(client)
	fs := frstr.NewFavoriteService(client)
	r := mux.NewRouter()

	s := server.NewServer(r, ws, us, fs)

	s.Run()
}
