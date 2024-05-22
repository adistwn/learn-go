package controllers

import (
	"net/http"

	"github.com/adistwn/learn-go/database/initializers"
	"github.com/adistwn/learn-go/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
