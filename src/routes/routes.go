package routes

import (
	"admin/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
}
