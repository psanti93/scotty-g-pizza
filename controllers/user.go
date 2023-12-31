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
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")

		user, err := uc.UserService.Create(email, password)

		if err != nil {
			http.Error(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}
		session, err := uc.SessionService.NewSession(user.ID)

		if err != nil {
			http.Error(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}

		setCookie(c.Writer, c.Request, CookieSession, session.Token.Session_Token)

		c.JSON(http.StatusOK, "User was created: "+user.Email)
	}

}

func (uc *UserController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc.Page.SignIn.Execute(c.Writer, c.Request, nil)
	}
}

func (uc *UserController) ProcessSignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")

		user, err := uc.UserService.Authenticate(email, password)

		if err != nil {
			http.Error(c.Writer, "Something went wrong: User doesn't exist", http.StatusInternalServerError)
			return
		}

		session, err := uc.SessionService.NewSession(user.ID)

		if err != nil {
			http.Error(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}

		setCookie(c.Writer, c.Request, CookieSession, session.Token.Session_Token)

		c.JSON(http.StatusOK, "User authenticated: "+user.Email)
	}

}

func (uc *UserController) CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := readCookie(c.Request, CookieSession)
		if err != nil {
			http.Error(c.Writer, "Unable to read cookie", http.StatusNotFound)
			return
		}

		user, err := uc.SessionService.CurrentSession(token)

		if err != nil {
			http.Redirect(c.Writer, c.Request, "/signup", http.StatusFound)
			return
		}

		c.JSON(http.StatusOK, "Current User: "+user.Email)
	}
}
