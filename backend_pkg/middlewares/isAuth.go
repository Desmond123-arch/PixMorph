package middlewares

import (
	"net/http"
	"strings"

	"github.com/Desmond123-arch/PixMorph.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		token := strings.Split(header, " ")
		if header == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		claims, err := utils.VerifyToken(token[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized"})
			ctx.Abort()
			return
		}
		details, ok := claims.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}
		username := details["username"].(string)
		ctx.Set("username", username)
		ctx.Next()
	}
}
