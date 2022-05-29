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

func (u *MockedUserRepo) GetUserByUsername(username string) (*users.User, error) {

	if username == "unkown" {
		return nil, Unauthorized
	}

	if username == "NeuerNutzer" {
		return nil, users.UserNotFound
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

func TestAuth(t *testing.T) {
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

	app := setupMockedEnv()

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
			t.Logf("Test: %v failed", test.description)
			t.Fail()
		}
	}
}

func TestAuthWithInvalidBody(t *testing.T) {
	app := setupMockedEnv()
	req := httptest.NewRequest("POST", "/api/auth", nil)
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}
	resp, _ := app.Test(req, -1)

	if resp == nil || 400 != resp.StatusCode {
		t.Fail()
	}
}

func setupMockedEnv() *fiber.App {
	var repo users.UserRepository = &MockedUserRepo{}
	users.UserRepo = repo
	app := fiber.New()
	RegisterAuthEndpoints(app)

	key = []byte("supersecretkey")

	return app
}
