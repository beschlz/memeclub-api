package main

import (
	"github.com/beschlz/memeclub-api/version"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	version.RegisterVersion(app)

	log.Fatal(app.Listen(":9090"))
}
