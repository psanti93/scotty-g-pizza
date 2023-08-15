package controllers

import (
	"net/http"

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

func (uc *UserController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct {
			Email string
		}
		data.Email = c.Request.FormValue("email")
		uc.Page.SignUp.Execute(c.Writer, c.Request, nil)
	}
}

func (uc *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		user := service.User{
			Email:         email,
			Password_Hash: password,
		}
		c.JSON(http.StatusOK, user)
	}

}

func (uc *UserController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc.Page.SignIn.Execute(c.Writer, c.Request, nil)
	}
}

func (uc *UserController) ProcessSignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
	}

}

func (uc *UserController) CurrentUser(c *gin.Context) {

}
