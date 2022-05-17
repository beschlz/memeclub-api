package users

import "github.com/gofiber/fiber/v2"

type MessageResponse struct {
	Message string `json:"message"`
}

type User struct {
	Username  string `gorm:"column:username;primaryKey" json:"username"`
	Useremail string `gorm:"column:user_email" json:"user_email"`
	Password  string `gorm:"column:password" json:"-"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Usermail string `json:"user_email"`
	Password string `json:"password"`
}

func RegisterUserRoutes(app *fiber.App) {
	app.Get("/api/users/:username", getUserByUsername)
	app.Post("/api/users", createUser)
}

func getUserByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	user, err := GetUserBayName(username)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func createUser(ctx *fiber.Ctx) error {

	request := new(CreateUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).Send(nil)
	}

	user, err := CreateUser(request)

	if err == InvalidUserName || err == UserAlreadyExistsErr {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(MessageResponse{
				Message: err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(MessageResponse{Message: "Well... We fucked up. No idea why this didn't work."})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
