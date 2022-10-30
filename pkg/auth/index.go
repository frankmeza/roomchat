package auth

import "github.com/golang-jwt/jwt"

const (
	HASH_COST = 10
)

type jwtClaims struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	IsAdmin bool   `json:"admin"`
	jwt.StandardClaims
}
