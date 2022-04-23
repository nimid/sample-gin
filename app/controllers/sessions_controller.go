package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Session struct{}

func (s Session) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "sessions/index.tmpl", gin.H{
			"title": "Example Gin",
		})
	}
}

func (s Session) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "a" && password == "b" {
			c.Redirect(http.StatusFound, "/dashboard")
		} else {
			c.Redirect(http.StatusFound, "/sessions")
		}
	}
}
