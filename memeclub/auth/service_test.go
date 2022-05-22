package auth

import (
	"github.com/beschlz/memeclub-api/memeclub/users"
	"testing"
)

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
