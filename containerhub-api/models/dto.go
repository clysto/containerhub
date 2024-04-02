package models

import "github.com/golang-jwt/jwt/v5"

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ContainerCreateParam struct {
	Image string `json:"image"`
}

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}
