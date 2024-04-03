package api

import (
	"containerhub-api/global"
	"containerhub-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	param := models.LoginParam{}
	if err := ctx.BindJSON(&param); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.DB.Where("username = ?", param.Username).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	if user.Password != param.Password {
		ctx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	claims := models.Claims{
		Username: user.Username,
		UserID:   user.ID,
		Admin:    user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "containerhub",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(global.Config.JWTSecret))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"token": t, "user": user})
}

func Signup(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Admin = false
	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User created"})
}
