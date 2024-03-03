package lib

import (
	"crypto/rand"
	"errors"
)

func GenerateRandomString(strLength int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomStringResult := make([]byte, strLength)
	randomBytes := make([]byte, strLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", errors.New("unable to generate random string")
	}
	const charsetSize = len(charset)
	for index, value := range randomBytes {
		randomStringResult[index] = charset[int(value)%charsetSize]
	}

	return string(randomStringResult), nil
}
