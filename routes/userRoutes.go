package routes

import (
	"github.com/gofiber/fiber/v2"
	userCtrl "github.com/muling3/go-fiber-jwt/controllers"
	"github.com/muling3/go-fiber-jwt/middlewares"
)

func ConfigureUserRoutes(app *fiber.App) {
	app.Post("/login", userCtrl.Login)

	//protecting the routes below with a custom jwt middleware
	app.Use(middlewares.JwtMiddleware)

	app.Get("/me", userCtrl.GetMe)
	app.Get("/", userCtrl.HomeRoute)
	app.Get("/users", userCtrl.GetAllUsers)
	app.Get("/:id", userCtrl.GetUser)

	app.Post("/", userCtrl.CreateUser)
	app.Post("/logout", userCtrl.Logout)
	
	app.Delete("/:id", userCtrl.DeleteUser)
}