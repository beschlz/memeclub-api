package version

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"testing"
)

func TestRegisterVersion(t *testing.T) {
	app := fiber.New()
	RegisterVersion(app)
	req := httptest.NewRequest("GET", "/api/version", nil)

	resp, _ := app.Test(req, 1)

	if 200 != resp.StatusCode {
		t.Fail()
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil || string(b) != ApiVersion {
		t.Fail()
	}

}
