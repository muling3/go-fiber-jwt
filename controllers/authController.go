package controllers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/muling3/go-fiber-jwt/data"
)

const AccessSecret = "access_secret"

var logger = log.New(os.Stderr, "users-api ", log.LstdFlags)

func Login(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	logger.Printf("Username: %v, Password: %v", username, password)
	user, err := data.Login(username, password)

	if err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
		})
	}

	//create jwts
	now := time.Now()

	c := jwt.StandardClaims{
		Subject:   user.Username,
		Issuer:    "muling3",
		IssuedAt:  jwt.NewTime(float64(now.Unix())),
		ExpiresAt: jwt.NewTime(float64(now.Add(time.Minute * 4).Unix())),
	}

	claims := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name), c)

	token, err := claims.SignedString([]byte(AccessSecret))

	if err != nil {
		ctx.SendStatus(fiber.StatusForbidden)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid token",
		})
	}

	//store token in cookies
	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"ok":           true,
		"message":      "Login route",
		"access_token": token,
	})
}

func GetMe(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("access_token")

	if cookie == "" {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "No access token",
		})
	}

	logger.Printf("Token: %v \n", cookie)

	//get claims from the token
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	})

	if err != nil {
		ctx.SendStatus(fiber.StatusForbidden)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Invalid token",
		})
	}

	// logger.Printf("%+v", token)
	claims := token.Claims.(*jwt.StandardClaims)

	return ctx.JSON(fiber.Map{
		"ok":       true,
		"username": claims.Subject,
	})
}

func Logout(ctx *fiber.Ctx) error {
	if ctx.Cookies("access_token") == "" {
		ctx.SendStatus(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"ok":      false,
			"message": "Already signed out",
		})
	}

	//store token in cookies
	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"ok":      true,
		"message": "Logged out successfully",
	})
}
