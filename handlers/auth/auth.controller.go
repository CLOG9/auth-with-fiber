package handlers

import (
	"log"
	"os"
	"testfiber/config"
	shared "testfiber/shared/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginCtrl(c *fiber.Ctx) error {
	//validation
	validate := validator.New()
	body := new(UserLogin)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
	}

	// Validate request body
	if err := validate.Struct(body); err != nil {
		return c.Status(401).JSON(shared.GlobErrResp(err.Error()))
	}

	user, err := GetUserByEmail(body.Email)
	if err != nil {
		return c.Status(403).JSON(shared.GlobErrResp(err.Error()))
	}
	if user.Password != body.Password {
		return c.Status(401).
			JSON(shared.GlobErrResp("email or password incorrect"))
	}
	sess, err := config.Store.St.Get(c)
	if err != nil {
		log.Println(err)
	}
	sess.Set(os.Getenv("SESSION_KEY"), os.Getenv("SESSION_SUB_KEY")+user.Email)
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return c.Status(200).JSON(shared.GlobResp("authenticated"))
}

func RegisterCtrl(c *fiber.Ctx) error {
	// Validation
	validate := validator.New()

	// Parse request body into UserRegister struct
	body := new(UserRegister)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
	}

	// Validate request body
	if err := validate.Struct(body); err != nil {
		return c.Status(401).JSON(shared.GlobErrResp(err.Error()))
	}
	if err := createUser(body); err != nil {
		return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
	}
	return c.JSON(shared.GlobResp(body))
}

func LogoutCtrl(c *fiber.Ctx) error {
	sess, err := config.Store.St.Get(c)
	if err != nil {
		log.Println(err)
	}

	sess.Delete(os.Getenv("SESSION_KEY"))

	// Destroy session
	if err := sess.Destroy(); err != nil {
		panic(err)
	}

	return c.Status(200).JSON(shared.GlobResp("logged out"))
}
