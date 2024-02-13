package admin_route

import (
	admin_controller "testfiber/handlers/admin"

	"github.com/gofiber/fiber/v2"
)

func Admin_route(api fiber.Router) {
	admin := api.Group("/admin")

	admin.Get("/dashboard", admin_controller.AdminCtrl)

}
