package controller

import (
	"time"

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

	ctx.SetCookie("access_token", tokens.AccessToken, int(10*time.Minute/time.Second), "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", tokens.RefreshToken, int(7*24*time.Hour/time.Second), "/", "localhost", false, true)

	webResponse := utils.NewWebResponse("Login successful", nil)

	ctx.JSON(200, webResponse)
}

func (a *AuthControllerImpl) Logout(ctx *gin.Context) {
	var payload *token.Payload = ctx.MustGet("user").(*token.Payload)

	if err := a.AuthService.Logout(payload.Sub.String()); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	webResponse := utils.NewWebResponse("Logout successful", nil)

	ctx.JSON(200, webResponse)
}

func (a *AuthControllerImpl) Refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		errorhandler.HandleError(ctx, errorhandler.NewUnauthorizedError("Token not found"))
		return
	}

	token, err := a.AuthService.Refresh(refreshToken)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", token.AccessToken, int(10*time.Minute/time.Second), "/", "localhost", false, true)

	webResponse := utils.NewWebResponse("Refresh token successful", nil)

	ctx.JSON(200, webResponse)
}

func (a *AuthControllerImpl) CsrfToken(ctx *gin.Context) {
	csrfToken := a.AuthService.CsrfToken()

	ctx.SetCookie("csrf_token", csrfToken, int(10*time.Minute/time.Second), "/", "localhost", false, true)
	webResponse := utils.NewWebResponse("CSRF token successful", csrfToken)

	ctx.JSON(200, webResponse)
}
