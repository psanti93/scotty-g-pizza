package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/psanti93/scotty-g-pizza/service"
)

type UserController struct {
	Page           Pages
	UserService    *service.UserService
	SessionService *service.SessionService
}

type Pages struct {
	SignIn View
	SignUp View
}

func NewUserController(p Pages, us *service.UserService, ss *service.SessionService) *UserController {
	return &UserController{
		Page:           p,
		UserService:    us,
		SessionService: ss,
	}
}

func (uc *UserController) SignIn(c *gin.Context) {

}

func (uc *UserController) CreateUser(c *gin.Context) {

}
