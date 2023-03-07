package v1

import (
	"auth/vladmsnk/internal/usecase"
	"auth/vladmsnk/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, a usecase.Auth) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	api := handler.Group("/")
	{
		newAuthRoutes(api, a, l)
	}
}
