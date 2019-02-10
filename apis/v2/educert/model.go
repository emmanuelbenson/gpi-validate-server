package educert

import (
	"github.com/emmanuelbenson/gpi-validate-v2/utils"
	"github.com/jinzhu/gorm"

	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// EducationalCertificate struct
type EducationalCertificate struct {
	gorm.Model
	UserID    int    `validate:"required"`
	FirstName string `validate:"required"`
	OtherName string `json:"otherName"`
	LastName  string `validate:"required"`
	Type      string `validate:"required"` // Type of educational certificate (pre-graduate, graduate/post-graduate, professional)
	Title     string `validate:"required"`
	Status    string `gorm:"default:'NEW'"`
	Verified  *int   `gorm:"default:0"`
	Document  string `validate:"required"` // Document to be validated
}

// Validate validates the above struct
func (edCert EducationalCertificate) Validate() []utils.Error {
	validate = validator.New()
	errs := utils.Errors

	err := validate.Struct(edCert)

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
