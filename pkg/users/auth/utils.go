package auth

import (
	"os"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordString(plaintext string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(plaintext),
		HASH_COST,
	)

	if err != nil {
		return "", errata.CreateError(err, errata.ErrMessage{
			Text: "GeneratePasswordString bcrypt.GenerateFromPassword",
		})
	}

	return string(passwordHash), nil
}

type CheckPasswordHashParams struct {
	Hash     string
	Password string
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

type GenerateTokenStringParams struct {
	Password string
	Username string
}

func GenerateTokenString(params GenerateTokenStringParams) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		IsAdmin: true,
		Name:    params.Username + params.Password,
		UUID:    uuid.NewV4().String(),
	})

	tokenAsString, err := token.SignedString(
		[]byte(os.Getenv(constants.JWT_SECRET)),
	)

	if err != nil {
		return "", errata.CreateError(err, errata.ErrMessage{
			Text: "GenerateTokenString token.SignedString",
		})
	}

	return tokenAsString, nil
}
