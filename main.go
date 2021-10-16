package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/http/rest"
	"github.com/schoolboybru/location-service/internal/repositories/postgres"
	"github.com/schoolboybru/location-service/internal/services/logic"
)

func main() {

	repository, err := postgres.New()

	if err != nil {
		panic(err)
	}

	service := logic.New(repository)

	handler := rest.NewHandler(service)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		locationGroup := v1.Group("/location")
		{
			locationGroup.GET("/:id", handler.Get)
			locationGroup.POST("/addLocation", handler.Post)
			locationGroup.GET("/country/:id", handler.GetByCountry)
			locationGroup.GET("city/:id", handler.GetByCity)
		}
	}

	router.Run(":8000")
}
