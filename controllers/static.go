package controllers

import "github.com/gin-gonic/gin"

func StaticHandler(tpl View) gin.HandlerFunc {
	return func(c *gin.Context) {
		tpl.Execute(c.Writer, c.Request, nil)
	}
}
