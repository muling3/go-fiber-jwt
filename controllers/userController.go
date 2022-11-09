package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/muling3/go-fiber-jwt/data"
	"github.com/muling3/go-fiber-jwt/models"
)

func HomeRoute(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"app-name":       "Go-fiber-Jwt",
		"author":         "muling3",
		"date-developed": "Nov 9 2022",
	})
}

func GetAllUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ok":    true,
		"users": data.GetUsers(),
	})
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Id is required",
		})
	}

	//collect Locals
	// logger.Println(ctx.Locals("username"))

	i, _ := strconv.Atoi(id)

	user := data.GetUser(i)

	if user.Id == 0 {
		ctx.SendStatus(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "User not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":      true,
		"message": "User fetched successfully",
		"user":    user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	var usr models.User

	if err := ctx.BodyParser(&usr); err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	user := data.AddUser(usr)

	ctx.SendStatus(fiber.StatusCreated)
	return ctx.JSON(fiber.Map{
		"ok":      true,
		"message": "User created successfully",
		"user":    user,
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Id is required",
		})
	}

	i, _ := strconv.Atoi(id)

	if err := data.RemoveUser(i); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":      true,
		"message": "User deleted successfully",
	})
}
