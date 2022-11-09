package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/muling3/go-fiber-jwt/controllers"
)

func JwtMiddleware(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()

	//check if authorization is provided
	if !strings.Contains(headers["Authorization"], "Bearer ") {
		ctx.SendStatus(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Unauthorized",
		})
	}

	token := strings.Split(headers["Authorization"], " ")[1]

	tn, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(controllers.AccessSecret), nil
	})

	if err != nil {
		ctx.SendStatus(fiber.StatusForbidden)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Forbidden",
		})
	}

	claims := tn.Claims.(*jwt.StandardClaims)

	ctx.Locals("username", claims.Subject)
	return ctx.Next()
}
