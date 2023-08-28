package service

import (
	"database/sql"
	"fmt"
)

const (
	DefaultBytes = 32
)

type Session struct {
	ID     int
	UserID int
	Token  *Token
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
func (ss *SessionService) NewSession(userId int) (*Session, error) {

	sessionToken, err := GenerateToken(DefaultBytes)

	if err != nil {
		return nil, fmt.Errorf("Error generating token: %v", err)
	}

	var session = Session{
		UserID: userId,
		Token:  sessionToken,
	}

	row := ss.DB.QueryRow(`INSERT INTO sessions(user_id,token_hash) VALUES($1,$2)
			ON CONFLICT(user_id) DO UPDATE SET token_hash=$2 RETURNING id`, session.UserID, session.Token.Session_Token_Hash)

	err = row.Scan(&session.ID)

	if err != nil {
		return nil, fmt.Errorf("Inserting session error: %v", err)
	}

	return &session, nil
}

// CurrentSession authenticates a user's session whnever one signs
func (ss *SessionService) CurrentSession(token string) (*User, error) {

	token_hashed := HashToken(token)

	var user User

	row := ss.DB.QueryRow(`SELECT users.id, users.email FROM users 
		JOIN sessions ON users.id = sessions.user_id WHERE sessions.token_hash=$1`, token_hashed)

	err := row.Scan(&user.ID, &user.Email)

	if err != nil {
		return nil, fmt.Errorf("Issue verifying user session: %v", err)
	}

	return &user, nil
}
