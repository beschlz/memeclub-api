package auth

import (
	"github.com/beschlz/memeclub-api/src/users"
	"testing"
)

type MockedUserRepo struct {
	users.UserRepository
}

func (u *MockedUserRepo) GetUserByUsername(username string) (*users.User, error) {
	user := users.User{
		Username:  username,
		Useremail: "bendt@schulz-hamburg.de",
		Password:  "$2a$10$pquTH6C9lwWPl8ty9eTkguPwCfuKqxB3x4Q57mDE866SbqOWKxYEW",
	}

	return &user, nil
}

func (u *MockedUserRepo) CreateUser(user *users.User) error {
	return nil
}

func TestAuthorizeUser(t *testing.T) {
	var repo users.UserRepository = &MockedUserRepo{}
	users.UserRepo = repo

	wrongCreds := &Credentials{
		Username: "besch",
		Password: "falschesPassword",
	}

	var token, err = AuthorizeUser(wrongCreds)

	if err == nil || token != "" {
		t.Fail()
	}

	correctCreds := &Credentials{
		Username: "besch",
		Password: "AlleMeineEntchen",
	}

	token, err = AuthorizeUser(correctCreds)

	if err != nil || token == "" {
		t.Fail()
	}
}
