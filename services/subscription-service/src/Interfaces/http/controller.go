package http

import (
	"net/http"
	"subscription-service/src/Applications/usecase"
	"subscription-service/src/Domains/entities"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AddSubscriptionHandler *usecase.AddSubscription
}

func (h *Handler) AddSubscription(c *gin.Context) {
	var subs entities.Subscriptions

	err := c.ShouldBindBodyWithJSON(&subs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	subs_id, err := h.AddSubscriptionHandler.Execute(subs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambahkan subscriptions",
		"data": gin.H{
			"id": subs_id,
		},
	})
}
