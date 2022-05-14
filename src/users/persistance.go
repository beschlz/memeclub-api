package users

import (
	"fmt"
	"github.com/beschlz/memeclub-api/src/database"
	"log"
)

type UserRepository interface {
	getUserByUsername(username string) (*User, error)
}

type UserRepositoryImpl struct {
	UserRepository
}

func (u *UserRepositoryImpl) getUserByUsername(username string) (*User, error) {
	var user *User
	dbErr := database.DB.Where("username = ?", username).Find(user)

	if dbErr.Error != nil {
		log.Printf("Error trying to get user with username %v\n", username)
		log.Println(dbErr)

		return user, dbErr.Error
	}

	if user.Username == "" {
		return user, fmt.Errorf("could not find a user with username %v", username)
	}

	return user, nil
}
