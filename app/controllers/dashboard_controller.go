package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Dashboard struct{}

func (d Dashboard) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard/index.tmpl", gin.H{
			"title": "Dashboard",
		})
	}
}
