package main

import (
    "admin/src/database"
    "admin/src/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

const (
    dsn = "admin:admin@tcp(db:3306)/ambassador?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {

    database.Connect()

    database.AutoMigrate()

    app := fiber.New()

    routes.Setup(app)

    // Setting CORS
    app.Use(cors.New(cors.Config {
        AllowCredentials: true,
    }))

    app.Listen(":3000")
}
