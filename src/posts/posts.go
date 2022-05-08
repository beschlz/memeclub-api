package posts

import (
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	PostId      int64  `gorm:"column:post_id;primaryKey;autoIncrement:true" json:"post_id"`
	Creator     string `gorm:"column:creator" json:"creator"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	ImageUrl    string `gorm:"column:image_url" json:"image_url"`
}

type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       []byte
}

func RegisterPosts(app *fiber.App) {
	app.Get("/api/posts", getPosts)
	app.Post("/api/posts", createPost)
}

func getPosts(ctx *fiber.Ctx) error {
	posts, err := GetAllPosts()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(posts)
}

func createPost(ctx *fiber.Ctx) error {
	createPostRequest := new(CreatePostRequest)

	if err := ctx.BodyParser(&createPostRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).Send(nil)
	}

	post, err := CreatePost(createPostRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).Send(nil)
	}

	if post == nil {
		return ctx.Status(fiber.StatusInternalServerError).Send(nil)
	}

	return ctx.Status(200).JSON(post)
}
