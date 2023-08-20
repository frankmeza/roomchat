package auth

import (
	"os"

	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordString(plaintext string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(plaintext),
		HASH_COST,
	)

	if err != nil {
		return "", errata.CreateError(err, []string{
			"GeneratePasswordString GenerateFromPassword",
		})
	}

	return string(passwordHash), nil
}

func CheckPasswordHash(params CheckPasswordHashParams) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(params.Hash),
		[]byte(params.Password),
	)

	if err != nil {
		return err
	}

	return nil
}

func GenerateTokenString(params GenerateTokenStringParams) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		IsAdmin: true,
		Name:    params.Username + params.Password,
		UUID:    appUtils.CreateUuid(),
	})

	tokenAsString, err := token.SignedString(
		[]byte(os.Getenv(constants.JWT_SECRET)),
	)

	if err != nil {
		return "", errata.CreateError(err, []string{
			"GenerateTokenString SignedString",
		})
	}

	return tokenAsString, nil
}
