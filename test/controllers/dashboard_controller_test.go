package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"sample-gin/app/controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("../../app/views/**/*")
	r.GET("/dashboard", controllers.Dashboard{}.Index())

	return r
}

func TestDashboard(t *testing.T) {
	r := setupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dashboard", nil)
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Dashboard")
}
