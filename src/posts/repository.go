package posts

import (
	"fmt"
	"github.com/beschlz/memeclub-api/src/database"
	"log"
)

type postRepositoryInterface interface {
	GetAllPosts() (*[]Post, error)
	Save(post *Post) error
	GetById(postId int64) error
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

func (postRepository *PostRepository) GetById(postId int64) (*Post, error) {
	var post Post
	result := database.DB.First(&post, postId)

	var err error
	if result.Error != nil {
		err = fmt.Errorf("could not create post with values %+v", post)
	}

	return &post, err
}

func (p *PostRepository) DeleteById(postId int64) error {
	post, err := p.GetById(postId)

	if err != nil {
		return err
	}

	result := database.DB.Delete(post)

	if result.Error != nil {
		return fmt.Errorf("error deleting post with id %v", postId)
	}

	return nil
}
