package say_hey

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SayHey(c *fiber.Ctx) error {
	fmt.Println("middlwaaaaaaare")
	return c.Next()
}
