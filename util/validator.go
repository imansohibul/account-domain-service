package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("nik", validateNIK)
	validate.RegisterValidation("fullname", validateFullname)
}

func GetValidator() *validator.Validate {
	return validate
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
