package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, h *Handler) {
	router.POST("/register", h.Register)
	fmt.Println("user router")
}
