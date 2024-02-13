package admin_route

import (
	"testfiber/config"
	admin_controller "testfiber/handlers/admin"

	"github.com/gofiber/fiber/v2"
)

func Admin_route(api fiber.Router) {
	admin := api.Group(config.RouteEndpts.Admin) // admin

	admin.Get(config.RouteEndpts.Dashboard, admin_controller.AdminCtrl)

}
