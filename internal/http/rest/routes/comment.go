package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/schoolboybru/location-service/internal/http/rest"
)

func AddCommentRoutes(rg *gin.RouterGroup, handler rest.Handlers) {
    commentGroup := rg.Group("/comment")
        
    commentGroup.GET("/:id", handler.GetComments)
    commentGroup.POST("/addComment", handler.PostComment)
    
}
