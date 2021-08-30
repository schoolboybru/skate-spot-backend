package main

import (
	"github.com/schoolboybru/location-service/internal/adding"
	"github.com/schoolboybru/location-service/internal/http/rest"
)

func main() {

	//router := gin.Default()
	var adder adding.Service

	handler := rest.Handler(adder)

	handler.Run(":8000")

	//v1 := router.Group("v1")
	//{
	//	locationGroup := v1.Group("/location")
	//	{
	//		location := new(controllers.LocationController)
	//		locationGroup.GET("/:id", location.GetLocation)
	//		locationGroup.POST("/addLocation", location.AddLocation)
	//		locationGroup.GET("/country/:id", location.GetLocationsByCountry)
	//		locationGroup.GET("city/:id", location.GetLocationsByCity)
	//	}
	//}

}
