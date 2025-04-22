package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// GenerateSecureNumber generates a cryptographically secure random numeric account number of a given length
func GenerateSecureNumber(length int) (string, error) {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		// Generate a random number between 0 and 9 using crypto/rand
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}
		// Append the random digit to the builder
		sb.WriteString(num.String())
	}

	return sb.String(), nil
}
