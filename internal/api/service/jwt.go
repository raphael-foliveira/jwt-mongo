package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
	"github.com/raphael-foliveira/fiber-mongo/internal/cfg"
)

var jwtService = &Jwt{}

type Jwt struct{}

func (js *Jwt) generateToken(payload schemas.TokenPayload) (string, error) {
	payload.IssuedAt = time.Now()
	payload.Expires = time.Now().Add(time.Hour * 5)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   payload.Sub,
		"email": payload.Email,
		"exp":   payload.Expires,
		"iat":   payload.IssuedAt,
	})
	return token.SignedString([]byte(cfg.JwtSecret))
}

func (js *Jwt) validateToken(token string) (*schemas.TokenPayload, error) {
	payload := &schemas.TokenPayload{}
	_, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if payload.Expires.Before(time.Now()) {
		return nil, &schemas.ApiErr{
			Code:    401,
			Message: "token expired",
		}
	}
	return payload, nil
}
