/*
	package model keeps all structures that are

represented in our database
*/
package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// struct User contains all fields of table "users"
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassowrd string `json:"-"`
}

// checking user field for validation
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u, validation.Field(
			&u.Email,            // we're checking email field
			validation.Required, // this field is required, cant be empty
			is.Email,            // check if its an email
		),
		validation.Field(
			&u.Password,
			validation.By(RequiredIf(u.EncryptedPassowrd == "")), // validation with condition
			validation.Length(6, 100),
		),
	)
}

// will be called before creating a record
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassowrd = enc
		return nil
	}

	return errors.New("password is empty")
}

// clears information that should not be displayed in response
func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePasswowrd(password string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.EncryptedPassowrd),
		[]byte(password),
	) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
