package main

import (
	// "fmt"
	"log"
	// "net/http"

	// "github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/muling3/go-fiber-jwt/db"
	"github.com/muling3/go-fiber-jwt/routes"
)

func main() {
	db.DbConn()
	
	app := fiber.New(fiber.Config{
		AppName:       "Go-fiber-Jwt",
		CaseSensitive: true,
	})

	//allowed creadentials middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// //custom jwt middleware
	// app.Use(jwtMiddleware)

	routes.ConfigureUserRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
