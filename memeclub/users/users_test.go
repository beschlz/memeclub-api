package users

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

func setupUserTest() *fiber.App {
	UserRepo = &MockedUserRepo{}

	app := fiber.New()

	RegisterUserRoutes(app)
	return app
}

func TestCreateUserRoute(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		route        string
		method       string
		userReq      CreateUserRequest
	}{
		{
			description:  "POST to /api/users with valid data should create a user",
			expectedCode: 201,
			route:        "/api/users",
			method:       "POST",
			userReq: CreateUserRequest{
				Username: "NeuerNutzer",
				Usermail: "newuser@newuser.de",
				Password: "meinPassword",
			},
		},
		{
			description:  "POST to /api/users for a user that already exists should fail",
			expectedCode: 400,
			route:        "/api/users",
			method:       "POST",
			userReq: CreateUserRequest{
				Username: "BestehenderNutzer",
				Usermail: "bestehenderNutzer@bestehenderNutzer.de",
				Password: "meinPassword",
			},
		},
		{
			description:  "POST to /api/users for a user with only whitespaces in name should fail",
			expectedCode: 400,
			route:        "/api/users",
			method:       "POST",
			userReq: CreateUserRequest{
				Username: " ",
				Usermail: "bestehenderNutzer@bestehenderNutzer.de",
				Password: "meinPassword",
			},
		},
		{
			description:  "POST to /api/users for a user with only whitespaces in email should fail",
			expectedCode: 400,
			route:        "/api/users",
			method:       "POST",
			userReq: CreateUserRequest{
				Username: "BestehenderNutzer",
				Usermail: " ",
				Password: "meinPassword",
			},
		},
	}

	app := setupUserTest()

	for _, test := range tests {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(test.userReq)

		if err != nil {
			t.Fail()
		}

		req := httptest.NewRequest(test.method, test.route, &buf)
		req.Header = map[string][]string{
			"Content-Type": {"application/json"},
		}
		resp, _ := app.Test(req, -1)

		if resp.StatusCode != test.expectedCode {
			t.Logf(test.description)
			t.Logf("Excpeted StatusCode %v, got %v", test.expectedCode, resp.StatusCode)
			t.Fail()
		}
	}
}

func TestGetUserByUsernameRoute(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		route        string
		method       string
	}{
		{
			description:  "/api/users/:username should return 404 when a user is requested that doest not exist",
			expectedCode: 404,
			route:        "/api/users/notfound",
			method:       "GET",
		},
		{
			description:  "/api/users/:username should return user when it exists",
			expectedCode: 200,
			route:        "/api/users/besch",
			method:       "GET",
		},
	}

	app := setupUserTest()

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		resp, _ := app.Test(req, -1)

		if resp.StatusCode != test.expectedCode {
			t.Logf(test.description)
			t.Fail()
		}
	}
}
