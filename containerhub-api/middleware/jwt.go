package middleware

import (
	"containerhub-api/global"
	"containerhub-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "Authorization header required"})
		ctx.Abort()
		return
	}
	var user models.Claims
	token, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWTSecret), nil
	})
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if !token.Valid {
		ctx.JSON(401, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
