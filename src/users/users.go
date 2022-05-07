package users

import "github.com/gofiber/fiber/v2"

type User struct {
	Username  string `gorm:"column:username;primaryKey" json:"username"`
	Useremail string `gorm:"column:user_email" json:"user_email"`
}

func RegisterUserRoutes(app *fiber.App) {
	app.Get("/api/users/:username", getUserByUsername)
}

func getUserByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	user, err := dbGetUserbyUsername(username)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
