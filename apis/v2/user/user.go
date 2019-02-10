package user

import (
	"github.com/jinzhu/gorm"

	"github.com/emmanuelbenson/gpi-validate-v2/utils"
	validator "gopkg.in/go-playground/validator.v8"
)

// User struct
type User struct {
	gorm.Model
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Token    string `json:"token"`
}

// Validate validates the User struct
func (user User) Validate() []utils.Error {
	validate = validator.New()
	errs := utils.Errors

	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			field := err.StructField()
			er := err.ActualTag()

			err := utils.Error{}
			err[field] = er
			errs = append(errs, err)
		}
	}

	return errs
}
