package main

import (
	"errors"
	"net/mail"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	email string
}

func CreateEmail(email string) (*Email, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return &Email{email}, nil
}

func isValidEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil && emailRegex.MatchString(address)
}

func (email Email) String() string {
	return email.email
}

func (email Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return email.email == other.email
}
