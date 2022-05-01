package version

import (
	"github.com/gofiber/fiber/v2"
)

const ApiVersion = "0.1.0"

func RegisterVersion(app *fiber.App) {
	app.Get("/api/version", getVersion)
}

func getVersion(ctx *fiber.Ctx) error {
	return ctx.SendString(ApiVersion)
}
