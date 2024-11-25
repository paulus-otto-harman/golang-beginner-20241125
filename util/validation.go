package util

import (
	"fmt"
	indonesia "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/id"
	"go.uber.org/zap"
	"project/service"
	validationrule "project/util/validation"
	"reflect"
	"strings"
)

type Validation struct {
	Validate   *validator.Validate
	Logger     *zap.Logger
	Translator ut.Translator
	ErrorBag   []Error
}

func InitValidator(service service.Service, log *zap.Logger) (service.Service, *zap.Logger, Validation) {
	myValidator := validator.New()
	myValidator.RegisterTagNameFunc(JsonFieldName)
	registerCustomValidationRules(myValidator, service)

	return service, log, Validation{Validate: myValidator, Logger: log, Translator: registerTranslator(myValidator)}
}

func (v *Validation) ValidateStruct(input interface{}) interface{} {
	err := v.Validate.Struct(input)
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(v.Translator))
		v.ErrorBag = append(v.ErrorBag, Error{
			Field:   e.Field(),
			Message: translatedErr.Error(),
			Tag:     e.Tag(),
			Value:   e.Value(),
			Param:   e.Param(),
		})
	}
	return v.ErrorBag
}

func JsonFieldName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

type Error struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value"`
	Param   string      `json:"param"`
}

func (err *Error) ShouldHaveJsonParam(input interface{}) Error {
	if err.Tag == "eqfield" {
		if param, validationHasParam := reflect.TypeOf(input).FieldByName(err.Param); validationHasParam {
			err.Param = param.Tag.Get("json")
		}
	}
	return *err
}

func registerCustomValidationRules(validator *validator.Validate, service service.Service) {
	customValidation := validationrule.Init(service.Validation)

	validator.RegisterValidation("unique", customValidation.Rule.Unique)
	validator.RegisterValidation("exists", customValidation.Rule.Exists)
	validator.RegisterValidation("exists-for-user", customValidation.Rule.ExistsForUser)
	validator.RegisterValidation("not-empty-cart", customValidation.Rule.NotEmptyCart)
	validator.RegisterValidation("is-regular", customValidation.Rule.RegularShipment)
	//validator.RegisterValidation("exists-for-user", func() validator.Func {
	//	return func(fl validator.FieldLevel) bool {
	//		return false
	//	}
	//}())
	//
	//validator.RegisterValidation("is-regular", func() validator.Func {
	//	return func(fl validator.FieldLevel) bool {
	//		return false
	//	}
	//}())

}

func registerTranslator(myValidator *validator.Validate) ut.Translator {
	idn := indonesia.New()
	uni := ut.New(idn, idn)
	trans, _ := uni.GetTranslator("id")
	_ = id.RegisterDefaultTranslations(myValidator, trans)
	return trans
}
