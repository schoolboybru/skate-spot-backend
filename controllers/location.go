package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/db/cache"
	"github.com/schoolboybru/location-service/model"
)

type LocationController struct{}

func (l LocationController) GetLocation(c *gin.Context) {
	if c.Param("id") != "" {
		l, err := cache.GetLocation(c.Param("id"))

		if err != nil {
			println(err)
		}

		if l.City != "" {
			c.JSON(http.StatusOK, l)
			return
		}

		location, err := model.GetLocationById(c.Param("id"))

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
	err := cache.SetLocation(location)

	if err != nil {
		println(err)
	}

	c.JSON(http.StatusOK, location)
}

func (l LocationController) GetLocationsByCountry(c *gin.Context) {
	if c.Param("id") != "" {
		var country = c.Param("id")
		locations, err := model.GetAllLocationsFromCountry(country)

		if err != nil {
			println(err)
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}

func (l LocationController) GetLocationsByCity(c *gin.Context) {
	if c.Param("id") != "" {
		var city = c.Param("id")
		locations, err := model.GetAllLocationsFromCity(city)

		if err != nil {
			println(err)
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}
