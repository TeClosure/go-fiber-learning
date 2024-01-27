package routes

import (
	"admin/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    // GROUP
    api := app.Group("api")
    admin := api.Group("admin")

    // POST
    admin.Post("register", controllers.Register)
    admin.Post("login", controllers.Login)
    admin.Post("logout", controllers.Logout)

    // GET
    admin.Get("user", controllers.User)
}
