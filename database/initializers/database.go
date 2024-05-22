package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")

	DB, err = gorm.Open(postgres.Open("host="+dbHost+" user="+dbUser+" dbname="+dbName+" port="+dbPort+" sslmode="+dbSslMode), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
