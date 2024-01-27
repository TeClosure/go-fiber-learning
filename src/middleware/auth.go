package middleware

import (
	"strconv"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(ctx *fiber.Ctx) error {
	log.Println("Do Middleware")
	// get from cookie
	cookie := ctx.Cookies("jwt")

	// get token
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil || !token.Valid {
		ctx.Status(fiber.StatusUnauthorized) // 401
		return ctx.JSON(fiber.Map {
			"message": "Your login information is incorrect.",
		})
	}

	return ctx.Next()
}

func GetUserID(ctx *fiber.Ctx) (uint, error) {
	// get from cookie
	cookie := ctx.Cookies("jwt")

	// get token
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}