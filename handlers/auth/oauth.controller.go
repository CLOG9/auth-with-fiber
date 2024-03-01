package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testfiber/config"
	shared "testfiber/shared/models"

	"io"

	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(
		os.Getenv("GOOGLE_STATE"),
	)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != os.Getenv("GOOGLE_STATE") {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken,
	)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userDataa, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	var userData UserData
	json.Unmarshal([]byte(userDataa), &userData)

	user, err := GetUserByEmail(userData.Email)
	if err != nil {
		userg := UserByEmail{
			Username: userData.Name,
			Email:    userData.Email,
			Password: "/",
			Type:     "OAUTH",
			Image:    userData.Picture,
			Token_id: userData.ID,
		}
		if err := createUser(&userg); err != nil {
			return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
		}
	}
	if user.Type != "OAUTH" {
		return c.Status(401).JSON(shared.GlobErrResp("login failed"))
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
