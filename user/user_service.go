package user

import (
	"errors"
	"user-service/database"
	utils "user-service/utlis"
)

type userService struct {
	Repo database.UserRepo
}

func NewUserService(repo database.UserRepo) UserInterface {
	return &userService{Repo: repo}
}

func (s *userService) SignUp(input UserInput) error {
	// Check if user already exists
	_, err := s.Repo.FindUserByEmail(input.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Encrypt password and create user
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}
	user := database.User{
		Email:     input.Email,
		Password:  hashedPassword,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	return s.Repo.CreateUser(&user)
}

// SignIn authenticates a user by email and password
func (s *userService) SignIn(input UserInput) (*User, error) {
	// Fetch the user by email
	user, err := s.Repo.FindUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Verify the password
	if !utils.CheckPassword(input.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Return user details
	return toUserModel(user), nil
}

func toUserModel(user *database.User) *User {
	return &User{
		Id:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (s *userService) ListUsers() ([]*User, error) {
	return s.toUserModelList(s.Repo.GetAllUsers())
}

func (s *userService) toUserModelList(users []database.User, err error) ([]*User, error) {
	if err != nil {
		return nil, err
	}
	result := make([]*User, len(users))
	for i, user := range users {
		result[i] = toUserModel(&user)
	}
	return result, nil
}
