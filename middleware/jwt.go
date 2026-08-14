package middleware

import (
	"net/http"
	"strings"

	"order-management/auth"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey   = "userID"
	ContextUsernameKey = "username"
	ContextRoleKey     = "role"
)

// JWTAuth validates the Bearer token on every request and injects the
// authenticated user's identity into the gin context.
func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		claims, err := auth.ParseToken(secret, parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextUsernameKey, claims.Username)
		c.Set(ContextRoleKey, claims.RoleName)

		c.Next()
	}
}
