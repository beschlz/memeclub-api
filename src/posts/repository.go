package posts

import (
	"fmt"
	"github.com/beschlz/memeclub-api/src/database"
	"log"
)

type postRepositoryInterface interface {
	GetAllPosts() (*[]Post, error)
	Save(post *Post) error
}

type PostRepository struct {
	postRepositoryInterface
}

func (postRepository *PostRepository) GetAllPosts() (*[]Post, error) {
	var posts = []Post{}
	result := database.DB.Find(&posts)

	if result.Error != nil {
		log.Println("Error accessing db")
		return &posts, result.Error
	}

	return &posts, nil
}

func (postRepository *PostRepository) Save(post *Post) error {
	result := database.DB.Save(post)

	var err error
	if result.Error != nil {
		err = fmt.Errorf("could not create post with values %+v", post)
	}

	return err
}
