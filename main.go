package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/schoolboybru/location-service/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

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

	router.Run(":8080")
}
