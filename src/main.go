package main

import (
	"github.com/beschlz/memeclub-api/src/database"
	"github.com/beschlz/memeclub-api/src/users"
	"github.com/beschlz/memeclub-api/src/version"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	database.InitDatabase()

	version.RegisterVersion(app)
	users.RegisterUserRoutes(app)

	log.Fatal(app.Listen(":9090"))
}
