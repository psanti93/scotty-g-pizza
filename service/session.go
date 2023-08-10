package service

import "database/sql"

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func NewSessionService(db *sql.DB) *SessionService {
	return &SessionService{
		DB: db,
	}
}

// NewSession creates a new session whenever a user signs in
func (ss *SessionService) NewSession(userId int) {
	//TODO: Save session whenever we sign in

}

// CurrentSession authenticates a user's session whnever one signs
func (ss *SessionService) CurrentSession(token string) {
	//TODO: Take the token from the session and validate by hashing it
}
