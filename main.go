package main

import (
	"github.com/gin-gonic/gin"
	"sample-gin/config"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("app/views/**/*")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	config.Routes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
