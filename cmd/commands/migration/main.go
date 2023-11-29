package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment/gormcomment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history/gormhistory"
	"github.com/mrbryside/rbh/domain/user/domain/user/gormuser"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Run()
	fmt.Println("Migration completed!!")
}

func Run() {
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		panic(err.Error())
	}
	defer close()

	db.AutoMigrate(&gormuser.User{})
	db.AutoMigrate(&gormappointment.Appointment{})
	db.AutoMigrate(&gormhistory.History{})
	db.AutoMigrate(&gormcomment.Comment{})
}
