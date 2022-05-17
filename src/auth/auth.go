package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var key = []byte(os.Getenv("JWT_SECRET"))

type TokenResponse struct {
	Token string `json:"token"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func RegisterAuthEndpoints(app *fiber.App) {
	app.Post("/api/auth", Auth)
}

func Auth(ctx *fiber.Ctx) error {
	var credentials Credentials

	err := ctx.BodyParser(&credentials)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).Send(nil)
	}

	token, err := AuthorizeUser(&credentials)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).Send(nil)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 30),
	})

	res := TokenResponse{Token: token}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
