package users

import (
	"fmt"
	"testing"
)

type MockedUserRepo struct {
	UserRepository
}

func (m *MockedUserRepo) getUserByUsername(username string) (*User, error) {

	if username == "notfound" {
		return nil, fmt.Errorf("User not found\n")
	}

	user := User{}

	user.Username = username
	user.Useremail = "user@mail.de"

	return &user, nil
}

func TestGetUserBayName(t *testing.T) {
	userRepo = &MockedUserRepo{}
	user, _ := GetUserBayName("besch")

	if user.Username != "besch" {
		t.Fatalf("Invalid")
	}

	_, err := GetUserBayName("notfound")

	if err == nil {
		t.Fatalf("User should not have been found")
	}

}
