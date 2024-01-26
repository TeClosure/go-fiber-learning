package controllers

import (
    "admin/src/database"
    "admin/src/models"

    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
    var data map[string]string

    if err:= ctx.BodyParser(&data); err != nil {
        return err
    }

    if data["password"] != data["password_confirm"] {
        ctx.Status(fiber.StatusBadRequest) // 400
        return ctx.JSON(fiber.Map{
            "message": "Password is incorrect.",
        })
    }

    pwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

    user := models.User{
        FirstName:      data["first_name"],
        LastName:       data["last_name"],
        Email:          data["email"],
        Password:       pwd,
        IsAmbassdor:    false,
    }

    database.DB.Create(&user)

    return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
    var data map[string]string

    if err := ctx.BodyParser(&data); err != nil {
        return err
    }

    var user models.User
    database.DB.Where("email = ?", data["email"]).First(&user)

    if user.ID == 0 {
        ctx.Status(fiber.StatusBadRequest)
        return ctx.JSON(fiber.Map {
            "message": "User not found.",
        })
    }

    // Check Password
    err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
    if err != nil {
        ctx.Status(fiber.StatusBadRequest)
        return ctx.JSON(fiber.Map {
            "message": "There is an error in your password.",
        })
    }
    return ctx.JSON(user)
}