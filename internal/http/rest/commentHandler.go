package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/internal/domain"
)

type CommentHandler interface {
    GetComments(c *gin.Context)
    PostComment(c *gin.Context)
}

func (h *handler) PostComment(c *gin.Context) {
    comment := domain.Comment{}

    err := c.BindJSON(&comment)

    if err != nil {
        println(err.Error())
    }

    c.JSON(http.StatusOK, comment)

    err = h.services.AddComment(&comment)

    if err != nil {
        println(err)
    }
}

func (h *handler) GetComments(c *gin.Context) {
    if c.Param("id") != "" {
        comments, err := h.services.GetComments(c.Param("id"))

        if err != nil {
            println(err.Error())
        }

        c.JSON(http.StatusOK, comments)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
        c.Abort()
    }
    return
}
