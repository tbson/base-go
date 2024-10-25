package vldtutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"src/common/ctype"
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

func ValidateUpdatePayload[T any](c echo.Context, target T) (ctype.Dict, error) {
	localizer := localeutil.Get()
	var result ctype.Dict
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}

	// Reset the body so it can be read again if needed
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Unmarshal into a map to get the keys present in the payload
	var payloadMap ctype.Dict
	if err := json.Unmarshal(bodyBytes, &payloadMap); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidJSONPayload,
		})

		return result, errutil.New("", []string{msg})
	}

	convertedMap := ctype.Dict{}
	targetType := reflect.TypeOf(target)
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		structKey := field.Name
		jsonTag := field.Tag.Get("json")
		fieldName := strings.Split(jsonTag, ",")[0]
		if value, ok := payloadMap[fieldName]; ok {
			convertedMap[structKey] = value
		}
	}

	return convertedMap, nil
}

func ValidatePayload[T any](c echo.Context, target T) (T, error) {
	localizer := localeutil.Get()
	var result T
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}

	// Reset the body so it can be read again if needed
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Unmarshal into a map to get the keys present in the payload
	var payloadMap ctype.Dict
	if err := json.Unmarshal(bodyBytes, &payloadMap); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidJSONPayload,
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
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FieldRequired,
		})
		for _, field := range missingFields {
			error.Add(field, []string{msg})
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
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.BadJson,
				TemplateData: ctype.Dict{
					"Offset": syntaxError.Offset,
				},
			})
			return result, errutil.New("", []string{msg})

		case errors.As(err, &unmarshalTypeError):
			// Report the specific field that has an invalid type
			fieldName := unmarshalTypeError.Field
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.InvalidFieldValue,
				TemplateData: ctype.Dict{
					"FieldName": fieldName,
					"Type":      unmarshalTypeError.Type.String(),
					"Value":     unmarshalTypeError.Value,
				},
			})
			return result, errutil.New(fieldName, []string{msg})

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			unknownField := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.NotRecognizedField,
			})
			return result, errutil.New(strings.Trim(unknownField, "\""), []string{msg})

		case errors.Is(err, io.EOF):
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.EmptyRequestBody,
			})
			return result, errutil.New("", []string{msg})

		default:
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.FailToDecodeJSON,
			})
			return result, errutil.New("", []string{msg})
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
				var msg string
				switch fe.Tag() {
				case "required":
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.FieldRequired,
					})
				case "oneof":
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.MustBeOneOf,
						TemplateData: ctype.Dict{
							"Values": fe.Param(),
						},
					})
				default:
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.InvalidValue,
					})
				}

				// Append the error message to the field's error list
				error.Add(fieldName, []string{msg})
			}
		} else {
			// For other errors, return a general message
			return result, errutil.New("", []string{err.Error()})
		}
		return result, &error
	}
	return data, nil
}

func ValidateId(id string) int {
	if id == "" {
		return 0
	}
	if id, err := strconv.Atoi(id); err == nil {
		return id
	}
	return 0
}

func ValidateIds(ids string) []int {
	var idList []int
	if ids == "" {
		return idList
	}
	for _, id := range strings.Split(ids, ",") {
		if id, err := strconv.Atoi(id); err == nil {
			idList = append(idList, id)
		}
	}
	return idList
}
