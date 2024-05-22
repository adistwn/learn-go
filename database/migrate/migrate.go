package main

import (
	"fmt"

	"github.com/adistwn/learn-go/config"
	"github.com/adistwn/learn-go/database/initializers"
	"github.com/adistwn/learn-go/models"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("Migrating database...")
	if !initializers.DB.Migrator().HasTable(&models.User{}) {
		err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
		if err != nil {
			panic("failed to migrate database")
		}

		fmt.Println("Successfully migrated database")
	}
}
