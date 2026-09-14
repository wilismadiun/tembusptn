package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/usecase"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
)

type Handler struct {
	RegisterHandler *usecase.Register
}

func (h *Handler) Register(c *gin.Context) {
	fmt.Println("controller 1")
	var user entities.User

	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userSuccess, err := h.RegisterHandler.Execute(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		fmt.Println("controller 2")
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambahkan user",
		"data":    userSuccess,
	})
}
