package controller

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/params"
	"github.com/Osas997/go-portfolio/internal/domains/auth/service"
	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/Osas997/go-portfolio/internal/pkg/token"
	"github.com/Osas997/go-portfolio/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService service.AuthService, validate *validator.Validate) AuthController {
	return &AuthControllerImpl{authService, validate}
}

func (a *AuthControllerImpl) Login(ctx *gin.Context) {
	var authRequest params.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		errorhandler.HandleError(ctx, errorhandler.NewBadRequestError("Invalid JSON", err.Error()))
		return
	}

	if err := a.validate.Struct(authRequest); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	tokens, err := a.AuthService.Login(&authRequest)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Login successful", tokens)

	ctx.JSON(200, webResponse)
}

func (a *AuthControllerImpl) Logout(ctx *gin.Context) {
	var payload *token.Payload = ctx.MustGet("user").(*token.Payload)

	if err := a.AuthService.Logout(payload.Sub.String()); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Logout successful", nil)

	ctx.JSON(200, webResponse)
}

func (a *AuthControllerImpl) Refresh(ctx *gin.Context) {
	var refreshRequest params.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&refreshRequest); err != nil {
		errorhandler.HandleError(ctx, errorhandler.NewBadRequestError("Invalid JSON", err.Error()))
		return
	}

	if err := a.validate.Struct(refreshRequest); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	token, err := a.AuthService.Refresh(&refreshRequest)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Refresh token successful", token)

	ctx.JSON(200, webResponse)
}
