package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Token struct {
	Session_Token      string
	Session_Token_Hash string
}

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
func GenerateToken(n int) (*Token, error) {
	bytes, err := bytes(n)
	if err != nil {
		fmt.Errorf("Error in creating bytes")
		return nil, fmt.Errorf("Error in generating token")
	}

	token_session := base64.URLEncoding.EncodeToString(bytes)

	token_session_hash := HashToken(token_session)

	var token = Token{
		Session_Token:      token_session,
		Session_Token_Hash: token_session_hash,
	}

	return &token, nil

}

func HashToken(token string) string {
	token_hash := sha256.Sum256([]byte(token))

	return base64.URLEncoding.EncodeToString(token_hash[:])
}
