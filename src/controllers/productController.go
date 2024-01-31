package controllers

import (
    "strconv"

	"admin/src/database"
	"admin/src/models"

	"github.com/gofiber/fiber/v2"
)

func Products(ctx *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return ctx.JSON(products)
}

func CreateProducts(ctx *fiber.Ctx) error {
	var product models.Product

	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return ctx.JSON(product)
}

func GetProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var product models.Product
	product.ID = uint(id)

	database.DB.Find(&product)

	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	product := models.Product {
		ID: uint(id),
	}

	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	product := models.Product {
		ID: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}