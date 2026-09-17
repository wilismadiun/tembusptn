package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenValidator interface {
	ValidateToken(token string) (string, string, error)
}

type userContext struct {
	UserId string
	Role   string
}

func AuthenticationAdmin(tokenValidator TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Authorization header is missing",
				},
			)
			return
		}

		parts := strings.SplitN(authorization, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid authorization format",
				},
			)
			return
		}

		token := parts[1]

		userID, role, err := tokenValidator.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid or expired token",
				},
			)
			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Access denied: admin role required",
				},
			)
			return
		}

		c.Set("user_identify", userContext{
			UserId: userID,
			Role:   role,
		})

		c.Next()
	}
}
