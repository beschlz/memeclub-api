package auth

import (
	"bytes"
	"encoding/json"
	"github.com/beschlz/memeclub-api/memeclub/users"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

type MockedUserRepo struct {
	users.UserRepository
}

// Mocks UserRepository.GetByUserName() Given the username 'unkown' this mock returns nil
// as a user and Unauthorized error
// Otherwise: It will return a valid user with the given name, the mail bendt@schulz-hamburg.de
// and a HasedPassword resembling "AlleMeineEntchen"
func (u *MockedUserRepo) GetUserByUsername(username string) (*users.User, error) {

	if username == "unkown" {
		return nil, Unauthorized
	}

	user := users.User{
		Username:  username,
		Useremail: "bendt@schulz-hamburg.de",
		Password:  "$2a$10$pquTH6C9lwWPl8ty9eTkguPwCfuKqxB3x4Q57mDE866SbqOWKxYEW",
	}

	return &user, nil
}

func (u *MockedUserRepo) CreateUser(_ *users.User) error {
	return nil
}

func TestAuthorizeUser(t *testing.T) {
	var repo users.UserRepository = &MockedUserRepo{}
	users.UserRepo = repo

	tests := []struct {
		description string
		creds       *Credentials
		excpetedErr error
	}{
		{
			description: "Valid username but wrong password",
			creds: &Credentials{
				Username: "besch",
				Password: "falschesPassword",
			},
			excpetedErr: Unauthorized,
		},
		{
			description: "Valid username with valid password",
			creds: &Credentials{
				Username: "besch",
				Password: "AlleMeineEntchen",
			},
			excpetedErr: nil,
		},
		{
			description: "Invalid username",
			creds: &Credentials{
				Username: "unknown",
				Password: "irrelevant",
			},
			excpetedErr: Unauthorized,
		},
	}

	for _, test := range tests {
		token, err := AuthorizeUser(test.creds)

		if test.excpetedErr == nil && err != nil {

			if token == "" {
				t.Logf("Test %v failed. Expected valid token.", test.description)
				t.Fail()
			}

		}

		if err != test.excpetedErr {
			t.Logf("Test %v failed", test.description)
			t.Fail()
		}
	}

}

func TestValidateToken(t *testing.T) {
	var repo users.UserRepository = &MockedUserRepo{}
	users.UserRepo = repo

	correctCreds := &Credentials{
		Username: "besch",
		Password: "AlleMeineEntchen",
	}

	token, _ := AuthorizeUser(correctCreds)

	validateTokenErr := ValidateToken(token)

	if validateTokenErr != nil {
		t.Fail()
	}

	wrongCreds := &Credentials{
		Username: "besch",
		Password: "falschesPassword",
	}

	token, _ = AuthorizeUser(wrongCreds)

	validateTokenErr = ValidateToken(token)

	if validateTokenErr == nil {
		t.Fail()
	}
}

func TestAuth(t *testing.T) {
	var repo users.UserRepository = &MockedUserRepo{}
	users.UserRepo = repo

	tests := []struct {
		description  string
		expectedCode int
		route        string
		body         Credentials
		method       string
	}{
		{
			description:  "Call /api/auth with GET",
			expectedCode: 405,
			method:       "GET",
			route:        "/api/auth",
			body: Credentials{
				Username: "",
				Password: "",
			},
		},
		{
			description:  "Valid call against /api/auth with correct Password and UserName",
			expectedCode: 200,
			method:       "POST",
			route:        "/api/auth",
			body: Credentials{
				Username: "besch",
				Password: "AlleMeineEntchen",
			},
		},
		{
			description:  "Call /api/auth with invalid username and password",
			expectedCode: 401,
			method:       "POST",
			route:        "/api/auth",
			body: Credentials{
				Username: "unkown",
				Password: "DefinitvKeinRichtigesPassword",
			},
		},
	}

	app := fiber.New()
	RegisterAuthEndpoints(app)

	for _, test := range tests {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(test.body)

		if err != nil {
			t.Fail()
		}

		req := httptest.NewRequest(test.method, test.route, &buf)
		req.Header = map[string][]string{
			"Content-Type": {"application/json"},
		}
		resp, _ := app.Test(req, -1)

		if resp == nil || test.expectedCode != resp.StatusCode {
			t.Fatalf("Test: %v failed", test.description)
		}
	}
}
