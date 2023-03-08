package v1

import (
	"auth/vladmsnk/internal/middlewares"
	"auth/vladmsnk/internal/usecase"
	"auth/vladmsnk/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GreetRoutes struct {
	t usecase.Greet
	l logger.Interface
}

// newGreetRoutes
func newGreetRoutes(api *gin.RouterGroup, g usecase.Greet, l logger.Interface, signingKey string) {
	r := &GreetRoutes{g, l}

	secured := api.Group("/api/v1").Use(middlewares.Auth(signingKey))
	{
		secured.GET("/greet", r.greet)
	}
}

func (g *GreetRoutes) greet(ctx *gin.Context) {
	greet, err := g.t.Greet(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, greet)
}
