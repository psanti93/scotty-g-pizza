package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		fmt.Errorf("bytes: %v", err)
		return nil, fmt.Errorf("bytes: %v", err)

	}
	if nRead < n {
		fmt.Errorf("Not enough bytes read")
		return nil, fmt.Errorf("Not enough bytes read")
	}

	return b, nil
}

// Generate Token returns a random string using crypto/rand.
// n is the number of bytes to be used to generate the random string
func GenerateToken(n int) (string, error) {
	bytes, err := bytes(n)

	if err != nil {
		fmt.Errorf("Error in creating bytes")
		return "", fmt.Errorf("Error in generating token")
	}

	token := base64.URLEncoding.EncodeToString(bytes)

	return token, nil

}
