package main

import (
	"fmt"
	"os"

	"github.com/adistwn/learn-go/config"
	"github.com/adistwn/learn-go/database/initializers"
	"github.com/adistwn/learn-go/router"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("Initializing...")
	r := gin.Default()
	router.GetRouter(r)

	r.Run(":" + os.Getenv("PORT"))
	fmt.Println("Apps Running at :" + os.Getenv("PORT"))
}
