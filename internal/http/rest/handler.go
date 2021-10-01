package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/domain"
)

type LocationHandler interface {
	Get(c *gin.Context)
	GetByCity(c *gin.Context)
	GetByCountry(c *gin.Context)
	Post(c *gin.Context)
}

type handler struct {
	addingService domain.Service
}

func NewHandler(a domain.Service) LocationHandler {
	return &handler{addingService: a}
}

func (h *handler) Post(c *gin.Context) {
	location := domain.Location{}

	c.BindJSON(&location)

	c.JSON(http.StatusOK, location)

	err := h.addingService.AddLocation(&location)

	if err != nil {
		println(err)
	}
}

func (h *handler) Get(c *gin.Context) {

	if c.Param("id") != "" {
		//l, err := cache.GetLocation(c.Param("id"))

		//if err != nil {
		//	println(err)
		//}

		//if l.City != "" {
		//	c.JSON(http.StatusOK, l)
		//	return
		//}

		location, err := h.addingService.GetLocation(c.Param("id"))

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

func (h *handler) GetByCity(c *gin.Context) {

	if c.Param("id") != "" {
		var city = c.Param("id")
		locations, err := h.addingService.GetLocationsByCity(city)

		if err != nil {
			println(err)
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}

func (h *handler) GetByCountry(c *gin.Context) {

	if c.Param("id") != "" {
		var country = c.Param("id")
		locations, err := h.addingService.GetLocationsByCountry(country)

		if err != nil {
			println(err)
		}

		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}
