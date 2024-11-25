package validationrule

import (
	"github.com/go-playground/validator/v10"
	"project/service"
	"strings"
)

type Validation struct {
	Rule Rule
}

type Rule struct {
	ValidationService *service.ValidationService
}

func Init(validationService *service.ValidationService) Validation {
	return Validation{Rule{ValidationService: validationService}}
}

func (r Rule) Unique(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), "#")
	unique, err := r.ValidationService.IsUnique(params[0], params[1], fl.Field().String())
	if err != nil {
		return false
	}
	return unique
}

func (r Rule) Exists(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), "#")
	exists, err := r.ValidationService.Exists(params[0], params[1], fl.Field())
	if err != nil {
		return false
	}
	return exists
}

func (r Rule) ExistsForUser(fl validator.FieldLevel) bool {
	if customerToken, _, ok := fl.GetStructFieldOK(); ok {
		exists, err := r.ValidationService.ExistsForUser(customerToken.String(), int(fl.Field().Int()))
		if err != nil {
			return false
		}
		return exists
	}
	return false
}

func (r Rule) NotEmptyCart(fl validator.FieldLevel) bool {
	notEmpty, err := r.ValidationService.NotEmptyCart(fl.Field().String())
	if err != nil {
		return false
	}
	return notEmpty
}

func (r Rule) RegularShipment(fl validator.FieldLevel) bool {
	return fl.Field().String() == "regular"
}
