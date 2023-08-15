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
	//1. Create Token Hash

	//2. Create a new session object with USER_ID, TOKEN, and TOKEN HASH

	//3. Insert New Session into tb

	//4. UPDATE later with ON Conflict

}

// CurrentSession authenticates a user's session whnever one signs
func (ss *SessionService) CurrentSession(token string) {
	//TODO: Take the token from the session and validate by hashing it

	//1. Hash the token

	//2. Create a new user var user User

	//3 Join the user and session table and get the user id, password hash and email
}

// TODO create function to hash teh token sha256 encode
