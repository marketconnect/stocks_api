package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint64
	Username string
	Password string
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username: username,
		Password: string(hashedPassword),
	}

	return user, nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
