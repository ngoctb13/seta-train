package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	}
}
