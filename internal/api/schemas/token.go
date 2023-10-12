package schemas

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Sub      string    `json:"sub"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"iat"`
	Expires  time.Time `json:"exp"`
	jwt.RegisteredClaims
}

type ValidateToken struct {
	Token string `json:"token"`
}
