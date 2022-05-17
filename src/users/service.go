package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var UserRepo UserRepository = &UserRepositoryImpl{}

var (
	UserAlreadyExistsErr = errors.New("user already exists")
	InvalidUserName      = errors.New("invalid user name. username must at least be three chars long")
)

func GetUserBayName(username string) (*User, error) {
	return UserRepo.GetUserByUsername(username)
}

func CreateUser(createUserRequest *CreateUserRequest) (*User, error) {
	_, err := GetUserBayName(createUserRequest.Username)

	if err == nil {
		return nil, UserAlreadyExistsErr
	}

	genereatedPw, hashError := bcrypt.GenerateFromPassword(
		[]byte(createUserRequest.Password),
		10)

	if hashError != nil {
		return nil, hashError
	}

	if validateUsername(createUserRequest.Username) != nil {
		return nil, InvalidUserName
	}

	user := User{
		Username:  createUserRequest.Username,
		Useremail: createUserRequest.Usermail,
		Password:  string(genereatedPw),
	}

	dbErr := UserRepo.CreateUser(&user)
	if dbErr != nil {
		return nil, dbErr
	}

	return &user, nil
}

func validateUsername(username string) error {

	if len(username) < 3 {
		return fmt.Errorf("username must at least be three chars long")
	}

	return nil
}
