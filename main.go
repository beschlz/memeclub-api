package main

import (
	"github.com/beschlz/memeclub-api/database"
	"github.com/beschlz/memeclub-api/version"
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

	version.RegisterVersion(app)

	log.Fatal(app.Listen(":9090"))
}
