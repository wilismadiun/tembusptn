package http

import (
	"net/http"
	usersubscriptions "subscription-service/src/Applications/usecase/userSubscriptions"
	"subscription-service/src/Domains/entities"

	"github.com/gin-gonic/gin"
)

type UserSubsHandler struct {
	AddUserSubscriptionHandler *usersubscriptions.AddUserSubScripitions
}

func (h *UserSubsHandler) AddUserSubScripitions(c *gin.Context) {
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

	var payload entities.CreateSubscriptionRequest
	err := c.ShouldBindBodyWithJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if payload.DurationInDays <= 0 || payload.SubscriptionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Make sure all required fields have been filled out correctly",
		})
		return
	}

	userSubs, err := h.AddUserSubscriptionHandler.Execute(userCtx.UserId, payload.SubscriptionID, payload.DurationInDays)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"data":    userSubs,
	})
}
