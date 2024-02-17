package handlers

import (
	"log"
	"testfiber/config"
	"testfiber/schema"
)

func createUser(body *UserRegister) error {
	// Create a new User object
	user := &schema.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}

	// Create the user record in the database
	if err := config.DB.Db.Create(user).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
