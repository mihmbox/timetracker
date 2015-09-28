package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const NotFoundByEmailError = errors.New("User not found by email")

// Get User by email from DB
func (user *User) LoadByEmail(email string) error {
	connection, err := GetConnection()
	if err != nil {
		return err
	}

	connection.Where(&User{Email: "email"}).First(&user)
	// validate found user Email
	if len(user.Email) == 0 {
		return NotFoundByEmailError
	}

	return nil
}

func (user *User) HashPassword(password string) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hpass
}
