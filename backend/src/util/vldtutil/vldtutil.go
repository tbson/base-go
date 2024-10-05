package vldtutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"src/util/errutil"
	"src/util/localeutil"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
)

// CustomValidator implements the echo.Validator interface
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the input struct using the validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// getRequiredFields extracts the required fields from the struct tags
func GetRequiredFields(v interface{}) []string {
	var requiredFields []string

	val := reflect.ValueOf(v)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		validateTag := field.Tag.Get("validate")
		if strings.Contains(validateTag, "required") {
			// Get the json tag to obtain the field name in the payload
			jsonTag := field.Tag.Get("json")
			fieldName := strings.Split(jsonTag, ",")[0] // Remove options like omitempty
			requiredFields = append(requiredFields, fieldName)
		}
	}

	return requiredFields
}

func ValidatePayload[T any](c echo.Context, target T) (T, error) {
	localizer := localeutil.Get()
	var result T
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "cannot_read_request_body",
				Other: "Can not read request body",
			},
		})
		return result, errutil.New("", []string{msg})
	}

	// Reset the body so it can be read again if needed
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Unmarshal into a map to get the keys present in the payload
	var payloadMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &payloadMap); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "invalid_json_payload",
				Other: "Invalid JSON payload",
			},
		})

		return result, errutil.New("", []string{msg})
	}

	// Check for missing required fields
	var requiredFields = GetRequiredFields(target)
	var missingFields []string
	for _, field := range requiredFields {
		if _, ok := payloadMap[field]; !ok {
			missingFields = append(missingFields, field)
		}
	}

	if len(missingFields) > 0 {
		// Collect errors for missing fields
		error := errutil.CustomError{}
		for _, field := range missingFields {
			error.Add(field, []string{"This field is required"})
		}
		return result, &error
	}

	// Decode the body into the struct and catch specific errors
	var data T
	decoder := json.NewDecoder(bytes.NewBuffer(bodyBytes))
	decoder.DisallowUnknownFields() // Disallow unknown fields in the payload

	if err := decoder.Decode(&data); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return result, errutil.New("", []string{fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)})

		case errors.As(err, &unmarshalTypeError):
			// Report the specific field that has an invalid type
			fieldName := unmarshalTypeError.Field
			return result, errutil.New(fieldName, []string{fmt.Sprintf("Invalid value for field '%s'. Expected type %v but got %v", fieldName, unmarshalTypeError.Type, unmarshalTypeError.Value)})

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			unknownField := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return result, errutil.New(strings.Trim(unknownField, "\""), []string{"This field is not recognized."})

		case errors.Is(err, io.EOF):
			return result, errutil.New("", []string{"Request body must not be empty."})

		default:
			return result, errutil.New("", []string{"Failed to decode JSON"})
		}
	}

	// Validate the struct
	if err := c.Validate(&data); err != nil {
		// Map to collect messages per field
		error := errutil.CustomError{}
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				// Map struct field name to JSON field name
				fieldName := fe.Field()
				structField, _ := reflect.TypeOf(data).FieldByName(fe.StructField())
				jsonTag := structField.Tag.Get("json")
				if jsonTag != "" {
					fieldName = strings.Split(jsonTag, ",")[0]
				}

				// Customize the error message based on the validation tag
				var errorMsg string
				switch fe.Tag() {
				case "required":
					errorMsg = "This field is required."
				case "oneof":
					errorMsg = fmt.Sprintf("Must be one of: %s.", fe.Param())
				default:
					errorMsg = "Invalid value."
				}

				// Append the error message to the field's error list
				error.Add(fieldName, []string{errorMsg})
			}
		} else {
			// For other errors, return a general message
			return result, errutil.New("", []string{err.Error()})
		}
		return result, &error
	}
	return data, nil
}
