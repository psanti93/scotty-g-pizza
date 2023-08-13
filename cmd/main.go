package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/psanti93/scotty-g-pizza/controllers"
	"github.com/psanti93/scotty-g-pizza/db"
	"github.com/psanti93/scotty-g-pizza/pages"
	"github.com/psanti93/scotty-g-pizza/views"
)

func main() {
	r := gin.Default()
	tpl := views.Must(views.ParseFS(pages.FS, "home.gohtml", "base-page.gohtml"))
	r.GET("/", controllers.StaticHandler(tpl))

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

	r.Run()
}
