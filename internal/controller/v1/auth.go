package v1

import (
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/usecase"
	"auth/vladmsnk/internal/util"
	"auth/vladmsnk/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, a usecase.Auth, l logger.Interface) {
	r := &AuthRoutes{a, l}

	handler.GET("/login", r.login)
	handler.POST("/user/register", r.register)
}

func (a *AuthRoutes) login(ctx *gin.Context) {
	var userLoginRequest dto.UserLoginRequest

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	response, err := a.t.GenerateToken(ctx, userLoginRequest)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, util.ErrInvalidPassword):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (a *AuthRoutes) register(ctx *gin.Context) {
	var userRegisterRequest dto.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userRegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	response, err := a.t.CreateUser(ctx, userRegisterRequest)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrUserAlreadyExists):
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}
	ctx.JSON(http.StatusOK, response)
}
