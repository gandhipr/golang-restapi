package validators

import (
	errmsg "apiserver/messages"
	"apiserver/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// validate declared global for caching.
var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("title", title); err != nil {
		log.Fatal("Error registering title validator. Err: ", err)
	}
	if err := validate.RegisterValidation("version", version); err != nil {
		log.Fatal("Error registering title validator. Err: ", err)
	}

}

// customized validator for "title" field.
func title(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[ ,_/.a-zA-Z0-9-]*$`).MatchString(fl.Field().String())
}

// customized validator for "version" field.
func version(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[_/.a-zA-Z0-9-]*$`).MatchString(fl.Field().String())
}

func ValidateInput(metadata utils.Metadata) (bool, map[string][]string) {
	if err := validate.Struct(metadata); err != nil {
		errors := make(map[string][]string)

		reflected := reflect.ValueOf(metadata)

		for _, err := range err.(validator.ValidationErrors) {

			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			//If yaml tag doesn't exist, use lower case of name.
			if name = strings.Split(field.Tag.Get("yaml"), ",")[0]; name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], errmsg.RequiredError)
				break
			case "email":
				errors[name] = append(errors[name], errmsg.EmailError)
				break
			case "title":
				errors[name] = append(errors[name], errmsg.TitleError)
				break
			case "version":
				errors[name] = append(errors[name], errmsg.VersionError)
				break
			default:
				errors[name] = append(errors[name], errmsg.DefaultError)
				break
			}
		}
		return false, errors
	}

	return true, nil
}
