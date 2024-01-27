package controllers

import (
    "strconv"
    "time"

    "admin/src/database"
    "admin/src/models"

    "github.com/gofiber/fiber/v2"
    "github.com/dgrijalva/jwt-go"
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

    user := models.User{
        FirstName:      data["first_name"],
        LastName:       data["last_name"],
        Email:          data["email"],
        IsAmbassdor:    false,
    }

    user.SetPassword(data["password"])

    result := database.DB.Create(&user)
    if result.Error != nil {
        ctx.Status(fiber.StatusBadRequest)
        return ctx.JSON(fiber.Map {
            "message": "The Email is already registered",
        })
    }

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
    err := user.ComparePassword(data["password"])
    if err != nil {
        ctx.Status(fiber.StatusBadRequest)
        return ctx.JSON(fiber.Map {
            "message": "There is an error in your password.",
        })
    }

    // Token Issuance
    payload := jwt.StandardClaims {
        Subject:    strconv.Itoa(int(user.ID)),
        ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
    }

    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

    if err != nil {
        ctx.Status(fiber.StatusBadRequest)
        return ctx.JSON(fiber.Map {
            "message": "Your login information is incorrect.",
        })
    }

    // Save to Cookie
    cookie := fiber.Cookie {
        Name:       "jwt",
        Value:      token,
        Expires:    time.Now().Add(time.Hour * 24),
        HTTPOnly:   true,
    }

    ctx.Cookie(&cookie)

    return ctx.JSON(fiber.Map {
        "message": "Success for save to cookie.",
    })
}

func Logout(ctx *fiber.Ctx) error {
    // Clear cookie
    cookie := fiber.Cookie {
        Name:   "jwt",
        Value:  "",
        Expires: time.Now().Add(-time.Hour * 24),
        HTTPOnly: true,
    }

    ctx.Cookie(&cookie)
    return ctx.JSON(fiber.Map {
        "message": "success",
    })
}

func User(ctx *fiber.Ctx) error {
    cookie := ctx.Cookies("jwt")

    token, err := jwt.ParseWithClaims(
        cookie,
        &jwt.StandardClaims{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        },
    )

    if err != nil {
        ctx.Status(fiber.StatusUnauthorized) // 401
        return ctx.JSON(fiber.Map {
            "message": "Not authenticated.",
        })
    }

    // Get userid
    payload := token.Claims.(*jwt.StandardClaims)

    // Search User
    var user models.User
    database.DB.Where("id = ?", payload.Subject).First(&user)
    return ctx.JSON(user)
}
