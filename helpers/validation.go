package helpers

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(structObj interface{}) map[string]string {
	uni = ut.New(en.New())
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()

	// Using the english translation package
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	validationErrors := validate.Struct(structObj)

	if validationErrors == nil {
		return nil
	}

	errorsMap := make(map[string]string)

	// Parsing the error messages inside a map
	for _, err := range validationErrors.(validator.ValidationErrors) {
		errorsMap[err.Field()] = err.Translate(trans)
	}

	return errorsMap
}
