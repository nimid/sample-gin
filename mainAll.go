package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mainAll() {
	r := gin.New()
	r.LoadHTMLGlob("app/views/**/*")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sessions/index.tmpl", gin.H{
			"title": "Example Gin",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard/index.tmpl", gin.H{
			"title": "Dashboard",
		})
	})

	r.GET("/sessions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sessions/index.tmpl", gin.H{
			"title": "Example Gin",
		})
	})

	r.POST("/", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "a" && password == "b" {
			c.Redirect(http.StatusFound, "/dashboard")
		} else {
			c.Redirect(http.StatusFound, "/sessions")
		}
	})

	err := r.Run()
	if err != nil {
		return
	}
}
