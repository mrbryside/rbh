package main

import (
	"errors"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrbryside/rbh/domain/user/domain/user/gormuser"
	"github.com/mrbryside/rbh/domain/user/types/myrole"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// seed
	err = UserIntegrationTestSeeder()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("seeding completed!!")
}

func UserIntegrationTestSeeder() error {
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		panic(err.Error())
	}
	defer close()

	user := gormuser.User{
		Name:     "Bryan",
		Email:    "bryan@mail.com",
		Password: "123456",
		Role:     myrole.PeopleTeam,
	}
	if result := db.Create(&user); result.Error != nil {
		return errors.New("error seeding user")
	}
	return nil

}
