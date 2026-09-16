package driver

import (
	"subscription-service/commons/middleware"
	infrastructures "subscription-service/src/Infrastructures"
	"subscription-service/src/Infrastructures/security"
	"subscription-service/src/Interfaces/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(router *gin.Engine, db *gorm.DB) {
	authValidator := security.AuthenticationTokenJWT{}

	// handler
	handler := infrastructures.Container(db)
	authMiddleware := middleware.Authentication(&authValidator)

	api := router.Group("/api")
	api.Use(authMiddleware)
	http.SubscriptionRouter(api, handler)
}
