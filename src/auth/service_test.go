package auth

import "testing"

func TestAuthorizeUser(t *testing.T) {
	wrongCreds := &Credentials{
		Username: "besch",
		Password: "falschesPassword",
	}

	var token, err = AuthorizeUser(wrongCreds)

	if err == nil || token != "" {
		t.Fail()
	}

	correctCreds := &Credentials{
		Username: "user1",
		Password: "password1",
	}

	token, err = AuthorizeUser(correctCreds)

	if err != nil || token == "" {
		t.Fail()
	}
}
