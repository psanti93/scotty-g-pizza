package service

import "database/sql"

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (us *UserService) Create(email, password string) {

}

func (us *UserService) Authenticate(email, password string) {

}
