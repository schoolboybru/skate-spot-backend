package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/skate-spot/internal/http/rest"
)

func AddPostRoutes(rg *gin.RouterGroup, handler rest.Handlers) {
	posts := rg.Group("/post")

	posts.GET("/:id", handler.GetPost)
	posts.POST("/addPost", handler.SavePost)
}
