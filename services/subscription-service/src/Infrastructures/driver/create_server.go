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
	subscriptionhandler := infrastructures.SubscriptionsContainer(db)
	userSubscriptionhandler := infrastructures.UserSubsContsiner(db)
	authMiddleware := middleware.AuthenticationAdmin(&authValidator)

	// router
	// subscriptions
	router.GET("/subscriptions", subscriptionhandler.GetAllSubscriptions)

	subscription := router.Group("/subscriptions")
	subscription.Use(authMiddleware)
	subscription.POST("", subscriptionhandler.AddSubscription)

	// user subscriptions
	userSubs := router.Group("/user-subscriptions")
	userSubs.Use(authMiddleware)
	userSubs.POST("", userSubscriptionhandler.AddUserSubScripitions)
}
