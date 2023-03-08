package middlewares

import (
	"auth/vladmsnk/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth(signingKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		headerSplitted := strings.Split(authHeader, " ")
		if len(headerSplitted) != 2 {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if headerSplitted[0] != "Bearer" {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := util.ValidateToken(headerSplitted[1], signingKey)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
