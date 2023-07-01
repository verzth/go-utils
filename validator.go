package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type CustomValidator struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func (cv *CustomValidator) SetTranslation(locale string) *CustomValidator {
	enT := en.New()
	idT := id.New()
	uni := ut.New(enT, idT, enT)
	trans, found := uni.GetTranslator(locale)
	if !found {
		panic("translator not found")
	}
	cv.Translator = trans

	var err error
	switch locale {
	case "id":
		err = id_translations.RegisterDefaultTranslations(cv.Validate, cv.Translator)
	case "en":
		err = en_translations.RegisterDefaultTranslations(cv.Validate, cv.Translator)
	default:
		err = en_translations.RegisterDefaultTranslations(cv.Validate, cv.Translator)
	}
	if err != nil {
		panic("translator not found")
	}

	return cv
}

func (cv *CustomValidator) GetErrMsg(err error) string {
	var msg string
	for i, e := range err.(validator.ValidationErrors) {
		if i > 0 {
			msg += ";" // Error message separator
		}
		msg += e.Translate(cv.Translator)
	}

	return msg
}

type ErrField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (cv *CustomValidator) GetErrMsgField(err error) (errF []ErrField) {
	for _, e := range err.(validator.ValidationErrors) {
		errF = append(errF, ErrField{
			Field:   e.Field(),
			Message: e.Translate(cv.Translator),
		})
	}
	return
}

func IsRFC3339(field validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, field.Field().String())
	return err == nil
}

func ConvertErrMsg[R any](errF any) ([]*R, error) {
	var result []*R
	b, err := json.Marshal(&errF)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewValidator() *CustomValidator {
	locale := "en"
	cv := &CustomValidator{
		Validate: validator.New(),
	}
	cv = cv.SetTranslation(locale)
	cv.Validate.RegisterValidation("isRFC3339", IsRFC3339)

	return cv
}
