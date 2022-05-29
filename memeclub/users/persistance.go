package users

import (
	"errors"
	"github.com/beschlz/memeclub-api/memeclub/database"
	"log"
)

type UserRepository interface {
	GetUserByUsername(username string) (*User, error)
	CreateUser(*User) error
}

type UserRepositoryImpl struct {
	UserRepository
}

var UserNotFound = errors.New("UserNotFound")

func (u *UserRepositoryImpl) GetUserByUsername(username string) (*User, error) {
	user := new(User)
	dbErr := database.DB.Where("username = ?", username).Find(user)

	if dbErr.Error != nil {
		log.Printf("Error trying to get user with username %v\n", username)
		log.Println(dbErr)

		return user, dbErr.Error
	}

	if user.Username == "" {
		return user, UserNotFound
	}

	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(user *User) error {
	result := database.DB.Save(user)

	return result.Error
}
