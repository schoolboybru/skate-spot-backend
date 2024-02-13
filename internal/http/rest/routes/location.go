package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/skate-spot/internal/http/rest"
)

func AddLocationRoutes(rg *gin.RouterGroup, handler rest.Handlers) {
	locations := rg.Group("/location")

	locations.GET("/:id", handler.Get)
	locations.POST("/addLocation", handler.Post)
	locations.GET("/country/:id", handler.GetByCountry)
	locations.GET("city/:id", handler.GetByCity)
}
