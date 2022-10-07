package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `validate:"email,required" json:"email"`
	Password          string `validate:"omitempty,min=6,max=100" json:"password,omitempty"`
	EncryptedPassword string `validate:"required_without=Password" json:"-"`
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

func (u *User) Sanitize() {
	u.Password = ""
}

func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
