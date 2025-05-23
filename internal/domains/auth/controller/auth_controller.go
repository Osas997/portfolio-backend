package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
