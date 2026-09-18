package driver

import (
	"subscription-service/commons/middleware"
	infrastructures "subscription-service/src/Infrastructures"
	"subscription-service/src/Infrastructures/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(router *gin.Engine, db *gorm.DB) {
	authValidator := security.AuthenticationTokenJWT{}

	// handler
	handler := infrastructures.Container(db)
	authMiddleware := middleware.AuthenticationAdmin(&authValidator)

	// router
	router.GET("/subscriptions", handler.GetAllSubscriptions)

	api := router.Group("/subscriptions")
	api.Use(authMiddleware)
	api.POST("", handler.AddSubscription)
}
