package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/psanti93/scotty-g-pizza/controllers"
	"github.com/psanti93/scotty-g-pizza/db"
	"github.com/psanti93/scotty-g-pizza/pages"
	"github.com/psanti93/scotty-g-pizza/service"
	"github.com/psanti93/scotty-g-pizza/views"
)

func main() {
	r := gin.Default()
	tpl := views.Must(views.ParseFS(pages.FS, "home.gohtml", "base-page.gohtml"))
	r.GET("/", controllers.StaticHandler(tpl))

	// TODO have database configuration set up on init
	dbCfg := db.DefaultConfig()
	db, err := db.Open(dbCfg)

	if err != nil {
		fmt.Errorf("Something went wrong: %v", err)
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	signUpPage := views.Must(views.ParseFS(pages.FS, "signup.gohtml", "base-page.gohtml"))
	signInPage := views.Must(views.ParseFS(pages.FS, "signin.gohtml", "base-page.gohtml"))
	pages := controllers.Pages{
		SignIn: signInPage,
		SignUp: signUpPage,
	}
	us := service.NewUserService(db)
	ss := service.NewSessionService(db)

	uc := controllers.NewUserController(pages, us, ss)

	r.GET("/signup", uc.SignUp())
	r.POST("/signup", uc.CreateUser())

	r.GET("/signin", uc.SignIn())
	r.POST("/signin", uc.ProcessSignIn())
	//TODO create current user url

	r.Run(":8081")
}
