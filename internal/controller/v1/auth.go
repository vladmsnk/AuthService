package v1

import (
	"auth/vladmsnk/internal/dto"
	"auth/vladmsnk/internal/usecase"
	"auth/vladmsnk/pkg/logger"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, response)
}
