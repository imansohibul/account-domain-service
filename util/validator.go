package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("nik", validateNIK)
	Validate.RegisterValidation("fullname", validateFullname)
}

func validateNIK(fl validator.FieldLevel) bool {
	nik := fl.Field().String()

	// NIK should be exactly 16 digits
	match, _ := regexp.MatchString(`^\d{16}$`, nik)
	return match
}

func validateFullname(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	// Accepts letters (upper/lower), spaces, dots, and hyphens
	// Disallows digits or special characters
	match, _ := regexp.MatchString(`^[a-zA-Z\s\.\-']{3,100}$`, name)
	return match
}
