package model

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}
