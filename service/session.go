package service

import "database/sql"

type SessionService struct {
	DB *sql.DB
}

func NewSessionService(db *sql.DB) *SessionService {
	return &SessionService{
		DB: db,
	}
}

// Create a new session
func (ss *SessionService) NewSession(userId int) {
	//TODO: Save session whenever we sign in

}

// Authenticate User that's on
func (ss *SessionService) CurrentSession(token string) {
	//TODO: Take the token from the session and validate by hashing it
}
