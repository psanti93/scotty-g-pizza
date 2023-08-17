package service

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int
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
func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Errorf("encrpt password: %v", err)
		return nil, err
	}

	password = string(hashedBytes)

	row := us.DB.QueryRow(`INSERT INTO users(email,password_hash) VALUES($1,$2) RETURNING id;`, email, password)

	user := User{
		Email:         email,
		Password_Hash: password,
	}
	row.Scan(&user.ID)

	return &user, nil
}

// Authenticate validates an existing user that signs in
func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{
		Email: email,
	}

	row := us.DB.QueryRow(`SELECT id, password_hash FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.Password_Hash)

	if err != nil {
		return nil, fmt.Errorf("Authenticate user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("authenticate user: %v", err)
	}

	return &user, nil

}
