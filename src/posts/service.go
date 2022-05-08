package posts

import (
	"fmt"
	"log"
	"strings"
)

func GetAllPosts() (*[]Post, error) {
	repo := PostRepository{}
	return repo.GetAllPosts()
}

func CreatePost(post *CreatePostRequest) (*Post, error) {
	if err := validatePost(post); err != nil {
		return nil, err
	}

	log.Printf("Creating new Post %v\n", post)

	postToSave := Post{
		PostId:      0,
		Creator:     "besch", // ToDo: After Auth use the user requsting the creation of this post
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    "", // ToDo: save image beforehand and insert  url here
	}

	repo := PostRepository{}
	if err := repo.Save(&postToSave); err != nil {
		return nil, err
	}

	return &postToSave, nil
}

func validatePost(postRequest *CreatePostRequest) error {
	var err error

	if strings.TrimSpace(postRequest.Title) == "" {
		err = fmt.Errorf("title of a post request must not consist of only whitespace character")
	}

	return err
}
