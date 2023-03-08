package middlewares

import (
	"auth/vladmsnk/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(signingKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := util.ValidateToken(tokenString, signingKey)
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
