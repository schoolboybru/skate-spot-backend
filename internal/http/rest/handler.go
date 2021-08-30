package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/adding"
)

func Handler(a adding.Service) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("v1")
	{
		locationGroup := v1.Group("/location")
		{
			//locationGroup.GET("/:id", location.GetLocation)
			locationGroup.POST("/addLocation", addLocation(a))
			//locationGroup.GET("/country/:id", location.GetLocationsByCountry)
			//locationGroup.GET("city/:id", location.GetLocationsByCity)
		}
	}

	return router

}

func addLocation(s adding.Service) func(c *gin.Context) {

	var location adding.Location
	return func(c *gin.Context) {
		c.BindJSON(&location)

		c.JSON(http.StatusOK, location)

		err := s.AddLocation(location)

		if err != nil {
			panic(err)
		}
	}
}
