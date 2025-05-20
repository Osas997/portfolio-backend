package middleware

import (
	"os"
	"strings"

	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/Osas997/go-portfolio/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := extractToken(c)
		if err != nil {
			errorhandler.HandleError(c, err)
			return
		}

		secret := os.Getenv("JWT_SECRET")
		claims, err := token.VerifyToken(tokenString, secret)
		if err != nil {
			appErr := errorhandler.NewUnauthorizedError("Invalid or expired token")
			errorhandler.HandleError(c, appErr)
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		appErr := errorhandler.NewUnauthorizedError("Token not found")
		return "", appErr
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		appErr := errorhandler.NewUnauthorizedError("Token not found")
		return "", appErr
	}

	// Extract the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	return tokenString, nil
}
