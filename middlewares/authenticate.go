package middlewares

import (
	"os"
	"strings"
	handlers "testfiber/handlers/auth"
	shared "testfiber/shared/models"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	cookie := c.Cookies(os.Getenv("COOKIE_NAME"))

	rdssession, err := handlers.GetRedisSession(cookie)
	if err != nil {
		println(err)
	}
	if !strings.Contains(string(rdssession), os.Getenv("SESSION_SUB_KEY")) {
		return c.Status(401).JSON(shared.GlobErrResp("unauhorized"))
	} // Proceed to the next handler
	return c.Next()
}
