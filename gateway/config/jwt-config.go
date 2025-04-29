package config

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClass struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetJWTSerectKey() string {
	return getEnv("JWT_SERECT_KEY", "")
}
