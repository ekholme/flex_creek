package server

import (
	"context"
	"net/http"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//interface to handle Wod stuff
type WodHandler interface {
	CreateWod(c *gin.Context)
	GetAllWods(c *gin.Context)
	GetRandomWod(c *gin.Context)
	GetWodByID(c *gin.Context)
	UpdateWod(c *gin.Context)
	DeleteWod(c *gin.Context)
}

//and a struct that can implement these methods
//might add user stuff to here later?
type wodHandler struct {
	wodService flexcreek.WodService
}

//create a new instance of wodHandler
func NewWodHandler(ws flexcreek.WodService) WodHandler {
	return &wodHandler{
		wodService: ws,
	}
}

//register Wod routes
func (s *Server) registerWodRoutes() {
	//endpoints for various handlers
	s.router.GET("/all_wods", s.wh.GetAllWods)
	s.router.GET("/random_wod", s.wh.GetRandomWod)

	//i think this one is right, but will want to try it
	s.router.GET("/wod/:id", s.wh.GetWodByID)

	//not sure what endpoint this should actually live at.
	s.router.POST("/create_wod", s.wh.CreateWod)
}

//TODO
func (wh wodHandler) CreateWod(c *gin.Context) {
	ctx := context.Background()

	var w *flexcreek.Wod

	//will pull data from a form eventually; this is sort of a placeholder
	err := c.ShouldBindJSON(&w)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//need to create an ID bc the form we pull this from won't do this by default
	id := uuid.New().String()

	w.ID = id

	err = wh.wodService.CreateWod(ctx, w)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wod successfully added!"})
}

func (wh wodHandler) GetAllWods(c *gin.Context) {
	//start here prob
}

func (wh wodHandler) GetRandomWod(c *gin.Context) {

}

func (wh wodHandler) GetWodByID(c *gin.Context) {

}

func (wh wodHandler) UpdateWod(c *gin.Context) {

}

func (wh wodHandler) DeleteWod(c *gin.Context) {

}
