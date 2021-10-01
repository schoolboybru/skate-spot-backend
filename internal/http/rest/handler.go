package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/adding"
)

type LocationHandler interface {
	Post(c *gin.Context)
}

type handler struct {
	addingService adding.Service
}

func NewHandler(a adding.Service) LocationHandler {
	return &handler{addingService: a}
}

func (h *handler) Post(c *gin.Context) {
	location := adding.Location{}

	c.BindJSON(&location)

	c.JSON(http.StatusOK, location)

	err := h.addingService.AddLocation(&location)

	if err != nil {
		println(err)
	}
}
