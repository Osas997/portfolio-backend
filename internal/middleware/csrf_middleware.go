package middleware

import (
	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/gin-gonic/gin"
)

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfCookie, err := c.Cookie("csrf_token")
		if err != nil {
			errorhandler.HandleError(c, errorhandler.NewUnauthorizedError("CSRF token not found"))
			return
		}

		csrfHeader := c.GetHeader("X-CSRF-Token")
		if csrfHeader == "" || csrfHeader != csrfCookie {
			errorhandler.HandleError(c, errorhandler.NewUnauthorizedError("CSRF token mismatch"))
			return
		}

		c.Next()
	}
}
