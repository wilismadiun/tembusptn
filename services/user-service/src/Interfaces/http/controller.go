package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Applications/usecase"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
)

type Handler struct {
	RegisterHandler *usecase.Register
	Loginhandler    *usecase.Login
}

func (h *Handler) Register(c *gin.Context) {
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
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambahkan user",
		"data":    userSuccess,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var login entities.Login

	err := c.ShouldBindBodyWithJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := h.Loginhandler.Execute(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menambahkan user",
		"data":    token,
	})
}
