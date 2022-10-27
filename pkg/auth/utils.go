package auth

import "golang.org/x/crypto/bcrypt"

func generatePasswordString(plaintext string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(plaintext),
		HASH_COST,
	)

	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
