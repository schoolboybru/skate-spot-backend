package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/skate-spot/internal/domain"
)

type LocationHandler interface {
	Get(c *gin.Context)
	GetByCity(c *gin.Context)
	GetByCountry(c *gin.Context)
	Post(c *gin.Context)
}

func (h *handler) Post(c *gin.Context) {
	location := domain.Location{}

	err := c.BindJSON(&location)

	if err != nil {
		println(err.Error())
	}

	c.JSON(http.StatusOK, location)

	err = h.services.AddLocation(&location)

	if err != nil {
		println(err)
	}
}

func (h *handler) Get(c *gin.Context) {

	if c.Param("id") != "" {

		location, err := h.services.GetLocation(c.Param("id"))

		if err != nil {
			println(err.Error())
		}

		c.JSON(http.StatusOK, location)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
	}
	return
}

func (h *handler) GetByCity(c *gin.Context) {

	if c.Param("id") != "" {
		var city = c.Param("id")
		locations, err := h.services.GetLocationsByCity(city)

		if err != nil {
			println(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}

func (h *handler) GetByCountry(c *gin.Context) {

	if c.Param("id") != "" {
		var country = c.Param("id")
		locations, err := h.services.GetLocationsByCountry(country)

		if err != nil {
			println(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}
