package main

import (
    "admin/src/database"

    "github.com/gofiber/fiber/v2"
)

const (
    dsn = "admin:admin@tcp(db:3306)/ambassador?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
    
    database.Connect()

    database.AutoMigrate()

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("My name is Self Note")
        // return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}