package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/http/rest"
	"github.com/schoolboybru/location-service/internal/repositories/postgres"
	"github.com/schoolboybru/location-service/internal/services"
	"github.com/schoolboybru/location-service/internal/http/rest/routes"
)

func main() {

	repository, err := postgres.New()

	if err != nil {
		panic(err)
	}

	service := services.New(repository)

	handler := rest.NewHandler(service)

	router := gin.Default()

	v1 := router.Group("v1")
    routes.AddCommentRoutes(v1, handler)
    routes.AddLocationRoutes(v1, handler)

	router.Run(":8000")
}
