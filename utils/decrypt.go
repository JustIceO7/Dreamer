package utils

import (
	"encoding/base64"

	"github.com/Strum355/log"
)

// Decrypts base64 string into bytes
func Decrypt(encodedData string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		log.WithError(err).Error("Failed to decode string formatting incorrect.")
		return nil, err
	}
	return decodedData, nil
}
