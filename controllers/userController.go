package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/muling3/go-fiber-jwt/db"
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
	users, err := db.FetchUsers()

	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"ok":    false,
			"message": "Error fetching users",
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":    true,
		"users": users,
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

	user, err := db.GetUserById(i)
	if err != nil {
		ctx.SendStatus(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
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

	id, err := db.RegisterUser(usr)

	if err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Error creating user",
		})
	}

	ctx.SendStatus(fiber.StatusCreated)
	return ctx.JSON(fiber.Map{
		"ok":      true,
		"message": "User created successfully",
		"userId":  id,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Id is required",
		})
	}

	i, _ := strconv.Atoi(id)

	var usr models.User

	if err := ctx.BodyParser(&usr); err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	u, err := db.UpdateUser(i, usr)
	if err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	ctx.SendStatus(fiber.StatusAccepted)
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"user": u,
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

	err := db.RemoveUser(i)
	if err != nil {
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
