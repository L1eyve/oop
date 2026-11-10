package id

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateId() (string, error) {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("ошибка в генерации id: %w", err)
	}

	id := base64.RawURLEncoding.EncodeToString(bytes)
	return id, nil
}
