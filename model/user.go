package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string `valid:"email,required"`
	Password          string `valid:"required,length(6|100)"`
	EncryptedPassword string
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

func (u *User) Validate() error {
	res, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("Validation failed")
	}
	return nil
}

func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
