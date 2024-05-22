package router

import (
	"github.com/adistwn/learn-go/app/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter(r *gin.Engine) {
	r.GET("/", controllers.GetUsers)
	r.POST("/api/login", controllers.Login)
}
