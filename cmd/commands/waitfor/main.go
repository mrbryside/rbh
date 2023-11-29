package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	for {
		err := dialDatabase(os.Getenv("DB_URL"), &gorm.Config{})
		if err == nil {
			fmt.Println("Successfully dialed to the database. Exiting the loop.")
			break
		}

		fmt.Printf("Error dialing the database: %v\n", err)
		time.Sleep(2 * time.Second)
	}
}

func dialDatabase(dataSourceName string, config *gorm.Config) error {
	_, err := gorm.Open(mysql.Open(dataSourceName), config)
	if err != nil {
		return err
	}

	return nil
}
