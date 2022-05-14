package users

var userRepo UserRepository = &UserRepositoryImpl{}

func GetUserBayName(username string) (*User, error) {
	return userRepo.getUserByUsername(username)
}
