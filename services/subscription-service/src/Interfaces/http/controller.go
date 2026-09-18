package http

import (
	"net/http"
	"subscription-service/src/Applications/usecase"
	"subscription-service/src/Domains/entities"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AddSubscriptionHandler    *usecase.AddSubscription
	GetAllSubscriptionHandler *usecase.GetAllSubscriptions
}

func (h *Handler) AddSubscription(c *gin.Context) {
	value, exist := c.Get("user_identify")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: user context not found",
		})
		return
	}

	userCtx, ok := value.(entities.UserContext)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse user context",
		})
		return
	}

	if userCtx.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Access denied: admin role required",
		})
		return
	}

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

func (h *Handler) GetAllSubscriptions(c *gin.Context) {
	subscriptions, err := h.GetAllSubscriptionHandler.Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(subscriptions) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data not found",
			"data":    subscriptions,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    subscriptions,
	})
}
