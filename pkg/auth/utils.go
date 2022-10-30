package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordString(plaintext string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(plaintext),
		HASH_COST,
	)

	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash), []byte(password),
	)

	return err == nil
}

// func generateTokenString(username, password string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
// 		IsAdmin: true,
// 		Name:    username + password,
// 		UUID:    uuid.NewV4().String(),
// 	})

// 	tokenAsString, err := token.SignedString(
// 		[]byte(os.Getenv(cc.JWT_SECRET)),
// 	)

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenAsString, nil
// }
