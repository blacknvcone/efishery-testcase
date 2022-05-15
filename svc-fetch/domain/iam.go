package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type IAMUser struct {
	ID       uint64 `json:"_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type IAMSession struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	jwt.Claims
}

type IIamUseCase interface {
	AuthorizationHTTP() gin.HandlerFunc
	IsAdmin(jwt.MapClaims) bool
}
