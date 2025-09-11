package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS(ctx *gin.Context) {
	// Build allowed origins list from env or default to local dev
	allowed := os.Getenv("ALLOWED_ORIGINS")
	if strings.TrimSpace(allowed) == "" {
		allowed = "http://localhost:5173"
	}
	allowedOrigins := map[string]struct{}{}
	for _, o := range strings.Split(allowed, ",") {
		origin := strings.TrimSpace(o)
		if origin != "" {
			allowedOrigins[origin] = struct{}{}
		}
	}

	origin := ctx.GetHeader("Origin")
	if origin != "" {
		if _, ok := allowedOrigins[origin]; ok {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Vary", "Origin")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
	}

	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT,PATCH, DELETE, OPTIONS")
	ctx.Header("Access-Control-Max-Age", "86400")

	if ctx.Request.Method == "OPTIONS" {
		// For preflight, if origin is allowed we already set headers above
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
