package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//placeholder stuff for now

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello World"})
	})

	r.Run(":8080")

}
