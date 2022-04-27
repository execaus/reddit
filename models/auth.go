package models

import (
	"errors"
	"fmt"
)

const MinPasswordLength = 8
const MinLoginLength = 6

type InputSignUp struct {
	Login    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
	Email    string `binding:"required"`
}

func (i *InputSignUp) IsValid() error {
	if len(i.Login) < MinLoginLength {
		return errors.New(fmt.Sprintf("login min length %d symbols", MinPasswordLength))
	}

	if len(i.Password) < MinPasswordLength {
		return errors.New(fmt.Sprintf("password min length %d symbols", MinPasswordLength))
	}

	return nil
}

type OutputSignUp struct {
	Session string `json:"session"`
}

type InputSignIn struct {
	Login    string
	Password string
}

type OutputSignIn struct {
	Session string  `json:"session"`
	Account Account `json:"account"`
}
