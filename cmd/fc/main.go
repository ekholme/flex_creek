package main

import (
	"github.com/ekholme/flex_creek/frstr"
	"github.com/ekholme/flex_creek/server"
	"github.com/gin-gonic/gin"
)

func main() {

	//maybe better to not do all of this in main.go? but it doesn't seem too bad for now

	//create router
	router := gin.Default()

	//create wodservice & wodhandler
	ws := frstr.NewWodService()
	wh := server.NewWodHandler(ws)

	s := server.NewServer(router, wh)

	s.Run()

}
