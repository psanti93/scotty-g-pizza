package main

import (
	"github.com/gin-gonic/gin"
	"github.com/psanti93/scotty-g-pizza/controllers"
	"github.com/psanti93/scotty-g-pizza/pages"
	"github.com/psanti93/scotty-g-pizza/views"
)

func main() {
	r := gin.Default()
	tpl := views.Must(views.ParseFS(pages.FS, "home.gohtml", "base-page.gohtml"))
	r.GET("/", controllers.StaticHandler(tpl))

	r.Run()
}
