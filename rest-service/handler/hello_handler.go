package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/handler/models"
)

func (h *Handler) HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		hello := models.Hello{Message: "Hello, World!"}
		c.JSON(http.StatusOK, hello)
	}
}
