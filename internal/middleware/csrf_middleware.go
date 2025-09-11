package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// hanya cek untuk method state-changing
		if strings.ToUpper(c.Request.Method) == http.MethodGet {
			c.Next()
			return
		}

		csrfHeader := c.GetHeader("X-CSRF-Token")
		csrfCookie, err := c.Cookie("csrf_token")
		if err != nil || csrfHeader == "" || csrfHeader != csrfCookie {
			c.AbortWithStatusJSON(419, gin.H{
				"error": "Invalid or missing CSRF token",
			})
			return
		}

		c.Next()
	}
}
