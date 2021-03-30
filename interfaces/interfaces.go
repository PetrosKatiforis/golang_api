package interfaces

import (
	"pioyi/golang_api/entity"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	jwt.StandardClaims
}

type ContextData struct {
	Token string  `json:"token"`
	Data  *Claims `json:"data"`
}

type UserResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}
