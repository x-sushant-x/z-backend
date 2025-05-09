package requests

import (
	"errors"
	"net/mail"
)

type AuthRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a AuthRequest) Validate(isSignUp bool) error {
	if a.Email == "" {
		return errors.New("email can't be empty")
	}

	_, err := mail.ParseAddress(a.Email)

	if err != nil {
		return errors.New("email not in proper format")
	}

	if a.Password == "" {
		return errors.New("password can't be empty")
	}

	if isSignUp {
		if a.Name == "" {
			return errors.New("name can't be empty")
		}
	}

	return nil
}
