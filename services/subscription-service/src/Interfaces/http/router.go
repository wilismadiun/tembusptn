package http

import "github.com/gin-gonic/gin"

func SubscriptionRouter(router *gin.RouterGroup, h *Handler) {
	router.POST("/subscriptions", h.AddSubscription)
}
