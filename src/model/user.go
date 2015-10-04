package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"logger"
)

var NotFoundByEmailError = errors.New("User not found by email")

// Get User by email from DB. Returns error if User was not found
func (user *User) LoadByEmail() error {
	connection, err := GetConnection()
	if err != nil {
		return err
	}

	logger.Info.Printf("User LoadByEmail email: %v", user.Email)
	if err := connection.Where(&User{Email: user.Email}).First(user).Error; err != nil {
		// == gorm.RecordNotFound
		return NotFoundByEmailError
	}

	return nil
}

// Creates DB entry.
// 'Created' field will be updated to current time.
// 'Password' will be encrypted. Provide original password!
func (user *User) Store() error {
	connection, err := GetConnection()
	if err != nil {
		return err
	}

	user.Created = time.Now()
	if err := user.HashPassword(); err != nil {
		return err
	}

	return connection.Create(user).Error
}

// Hashes current User password by bcrypt lib.
func (user *User) HashPassword() error {
	hpass, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hpass
	return nil
}

// Compares User password with specified value using bcrypt package's secure comparison.
func (user *User) ValidatePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
