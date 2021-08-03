package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/model"
)

type LocationController struct{}

var locationModel = new(model.Location)

func (l LocationController) GetLocation(c *gin.Context) {
	if c.Param("id") != "" {
		location, err := model.GetById(c.Param("id"))

		if err != nil {
			println(err)
		}

		c.JSON(http.StatusOK, location)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
	}
	return
}

func (l LocationController) AddLocation(c *gin.Context) {
	var location model.Location
	c.BindJSON(&location)
	model.AddLocation(location)
	c.JSON(http.StatusOK, gin.H{"new location": location.Name})
}
