package gen

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateRandomString(size int) ([]byte, error) {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

func GenPasswordWithSalt(password, salt []byte) (string, error) {
	var sha512Hasher = sha256.New()

	// Append salt to gen
	password = append(password, salt...)

	// Write gen bytes to the hasher
	if _, err := sha512Hasher.Write(password); err != nil {
		return "", err
	}

	// Get the SHA-256 hashed gen
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	// Convert the hashed gen to a base64 encoded string
	base64EncodedPasswordHash :=
		base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash, nil
}
