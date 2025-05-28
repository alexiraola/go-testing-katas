package entities

import (
	"errors"
	"hexagonal/domain/valueObjects"
)

type User struct {
	password *valueObjects.Password
	id       *valueObjects.Id
	email    *valueObjects.Email
}

func CreateUser(password *valueObjects.Password, id *valueObjects.Id, email *valueObjects.Email) *User {
	return &User{password: password, id: id, email: email}
}

func (user *User) ChangePassword(newPassword *valueObjects.Password) error {
	if newPassword.String() == user.password.String() {
		return errors.New("new password must be different")
	}
	user.password = newPassword
	return nil
}

func (user *User) IsMatchingPassword(password *valueObjects.Password) bool {
	return *user.password == *password
}
