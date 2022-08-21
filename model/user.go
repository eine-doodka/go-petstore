package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string `validate:"email,required"`
	Password          string `validate:"omitempty,min=6,max=100"`
	EncryptedPassword string `validate:"required_without=Password"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func (u *User) Validate(v *validator.Validate) error {
	return v.Struct(u)
}

func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
