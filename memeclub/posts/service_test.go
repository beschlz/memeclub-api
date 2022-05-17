package posts

import (
	"testing"
)

type MockPostRepository struct {
	PostRepositoryInterface
}

func (m *MockPostRepository) GetAllPosts() (*[]Post, error) {
	posts := []Post{
		{
			PostId:      0,
			Creator:     "besch",
			Title:       "Eins",
			Description: "Erster Post",
			ImageUrl:    "minio://memes/ersterPost.jpg",
		},
		{
			PostId:      1,
			Creator:     "bendt",
			Title:       "Zwei",
			Description: "Zweiter Post",
			ImageUrl:    "minio://memes/zweiterPost.jpg",
		},
	}

	return &posts, nil
}

func (m *MockPostRepository) Save(post *Post) error {
	return nil
}

func (m *MockPostRepository) GetById(postId int64) (*Post, error) {
	return &Post{
		PostId:      postId,
		Creator:     "besch",
		Title:       "Mein Titel",
		Description: "Meine Beschreibung",
		ImageUrl:    "minio://memes/langweiligesMeme.jpg",
	}, nil
}

func (m *MockPostRepository) DeleteById(postId int64) error {
	return nil
}

func TestGetAllPosts(t *testing.T) {
	PostRepo = &MockPostRepository{}
	posts, err := GetAllPosts()

	if len(*posts) != 2 {
		t.Fatalf("Excepted to get 2 posts. Got %v\n", len(*posts))
	}

	if err != nil {
		t.Fatalf("Unexpected error. Got %v\n", err)
	}
}

func TestCreatePostWithInvalidRequestZeroValuesSet(t *testing.T) {
	PostRepo = &MockPostRepository{}
	postRequest := CreatePostRequest{
		Title:       "",
		Description: "",
		Image:       nil,
	}

	post, err := CreatePost(&postRequest)

	checkPostNotNilAndErrorNil(post, err, &postRequest, t)
}

func TestCreatePostWithNoTitle(t *testing.T) {
	PostRepo = &MockPostRepository{}
	postRequest := CreatePostRequest{
		Title:       "",
		Description: "Eine Beschreibung",
		Image:       nil,
	}

	post, err := CreatePost(&postRequest)

	checkPostNotNilAndErrorNil(post, err, &postRequest, t)
}

func TestCreatePostWithTitleAndDescription(t *testing.T) {
	PostRepo = &MockPostRepository{}
	postRequest := CreatePostRequest{
		Title:       "A",
		Description: "Eine Beschreibung",
		Image:       nil,
	}

	post, err := CreatePost(&postRequest)

	if post == nil || err != nil {
		t.Fatalf("Valid post was not created. Post request was %v\n", postRequest)
	}
}

func checkPostNotNilAndErrorNil(post *Post,
	err error,
	postRequest *CreatePostRequest,
	t *testing.T) {

	if post != nil || err == nil {
		t.Errorf("Created Post with invalid data. Created Post %v with Request %v\n",
			post,
			postRequest)
	}
}
