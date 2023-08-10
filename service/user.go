package service

import "database/sql"

type User struct {
	ID            int
	Name          string
	Email         string
	Password_Hash string
}

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

// Create inserts a new user into the database
func (us *UserService) Create(email, password string) {

}

// Authenticate validates an existing user that signs in
func (us *UserService) Authenticate(email, password string) {

}
