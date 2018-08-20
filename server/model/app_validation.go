package model

import (
	"gopkg.in/go-playground/validator.v9"
)

func (m *User) Validate() error {
	validator := validator.New()
	return validator.Struct(m)
}

func (m *Memo) Validate() error {
	validator := validator.New()
	return validator.Struct(m)
}
