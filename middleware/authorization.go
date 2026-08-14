package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Role names used throughout the application. Keeping them as constants
// avoids typos when wiring routes (see routes/routes.go).
const (
	RoleAdmin   = "ADMIN"
	RoleManager = "MANAGER"
	RoleUser    = "USER"
)

// RequireRoles restricts access to users whose role (set by JWTAuth) is
// included in the given whitelist. Must run after JWTAuth.
func RequireRoles(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(roles))
	for _, r := range roles {
		allowed[r] = true
	}

	return func(c *gin.Context) {
		roleVal, exists := c.Get(ContextRoleKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		role, ok := roleVal.(string)
		if !ok || !allowed[role] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}
