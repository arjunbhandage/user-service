package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
}

type UserRepo interface {
	CreateUser(user *User) error
	// FindUserByEmail fetches a user by their email
	FindUserByEmail(email string) (*User, error)
	// GetAllUsers retrieves all user records
	GetAllUsers() ([]User, error)
	// UpdateUser updates an existing user's details
	UpdateUser(user *User) error
	// DeleteUser removes a user record from the database by ID
	DeleteUser(userID uint) error
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo returns a new instance of UserRepo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindUserByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) DeleteUser(userID uint) error {
	return r.db.Delete(&User{}, userID).Error
}
