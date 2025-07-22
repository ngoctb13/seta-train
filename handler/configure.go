package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/internal/domains/user/usecases"
)

type Handler struct {
	user *usecases.User
}

func NewHandler(user *usecases.User) *Handler {
	return &Handler{
		user: user,
	}
}

func (h *Handler) ConfigAuthRouteAPI(router *gin.RouterGroup) {
}
