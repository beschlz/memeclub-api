package main

import (
	"github.com/beschlz/memeclub-api/memeclub/auth"
	"github.com/beschlz/memeclub-api/memeclub/database"
	"github.com/beschlz/memeclub-api/memeclub/posts"
	"github.com/beschlz/memeclub-api/memeclub/users"
	"github.com/beschlz/memeclub-api/memeclub/version"
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
	database.InitMinio()

	version.RegisterVersion(app)
	users.RegisterUserRoutes(app)
	posts.RegisterPosts(app)
	auth.RegisterAuthEndpoints(app)

	log.Fatal(app.Listen(":9090"))
}
