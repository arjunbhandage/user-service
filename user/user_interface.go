package user

type UserInterface interface {
	SignUp(input UserInput) error
	SignIn(input UserInput) (*User, error)
	ListUsers() ([]*User, error)
}
