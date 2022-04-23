package config

import (
	"github.com/gin-gonic/gin"
	"sample-gin/app/controllers"
)

func Routes(r *gin.Engine) {
	r.GET("/", controllers.Session{}.Index())
	r.GET("/dashboard", controllers.Dashboard{}.Index())
	r.GET("/sessions", controllers.Session{}.Index())
	r.POST("/sessions", controllers.Session{}.Create())
}
