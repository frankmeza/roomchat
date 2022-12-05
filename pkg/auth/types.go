package auth

import "github.com/golang-jwt/jwt"

const (
	HASH_COST = 10
)

type (
	CheckPasswordHashParams struct {
		Hash     string
		Password string
	}

	GenerateTokenStringParams struct {
		Password string
		Username string
	}

	JwtClaims struct {
		Name    string `json:"name"`
		UUID    string `json:"uuid"`
		IsAdmin bool   `json:"admin"`
		jwt.StandardClaims
	}
)
