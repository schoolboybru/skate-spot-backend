package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/controllers"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("v1")
	{
		locationGroup := v1.Group("/location")
		{
			location := new(controllers.LocationController)
			locationGroup.GET("/:id", location.GetLocation)
			locationGroup.POST("/addLocation", location.AddLocation)
			locationGroup.GET("/country/:id", location.GetLocationsByCountry)
			locationGroup.GET("city/:id", location.GetLocationsByCity)
		}
	}

	router.Run(":8000")
}
