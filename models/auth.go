package models

import (
	"errors"
	"fmt"
	"strings"
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

	if strings.Contains(i.Login, "@") {
		return errors.New(`login have forbidden symbol "@"`)
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
	Identifier string
	Password   string
}

func (i *InputSignIn) Validate() error {
	if i.Identifier == "" {
		return errors.New("login required")
	}

	if i.Password == "" {
		return errors.New("password required")
	}

	return nil
}

type OutputSignIn struct {
	Session string  `json:"session"`
	Account Account `json:"account"`
}
