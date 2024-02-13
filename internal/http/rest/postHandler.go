package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/skate-spot/internal/domain"
)

type PostHandler interface {
	GetPost(c *gin.Context)
	SavePost(c *gin.Context)
}

func (h *handler) SavePost(c *gin.Context) {
	post := domain.Post{}

	err := c.BindJSON(&post)

	if err != nil {
		println(err.Error())
	}

	c.JSON(http.StatusOK, post)

	err = h.services.AddPost(&post)

	if err != nil {
		println(err)
	}
}

func (h *handler) GetPost(c *gin.Context) {
	if c.Param("id") != "" {
		post, err := h.services.GetPost(c.Param("id"))

		if err != nil {
			println(err.Error())
		}

		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
	}
	return
}
