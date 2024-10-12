package localeutil

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	CannotReadRequestBody = &i18n.Message{
		ID:    "CannotReadRequestBody",
		Other: "Can not read request body",
	}

	InvalidJSONPayload = &i18n.Message{
		ID:    "InvalidJSONPayload",
		Other: "Invalid JSON payload",
	}

	FieldRequired = &i18n.Message{
		ID:    "FieldRequired",
		Other: "This field is required",
	}

	BadJson = &i18n.Message{
		ID:    "BadJson",
		Other: "Request body contains badly-formed JSON (at position {{.Offset}})",
	}

	InvalidFieldValue = &i18n.Message{
		ID:    "InvalidFieldValue",
		Other: "Invalid value for field '{{.FieldName}}'. Expected type {{.Type}} but got {{.Value}}",
	}

	InvalidValue = &i18n.Message{
		ID:    "InvalidValue",
		Other: "Invalid value",
	}

	NotRecognizedField = &i18n.Message{
		ID:    "NotRecognizedField",
		Other: "This field is not recognized",
	}

	FailToDecodeJSON = &i18n.Message{
		ID:    "FailToDecodeJSON",
		Other: "Failed to decode JSON",
	}

	EmptyRequestBody = &i18n.Message{
		ID:    "EmptyRequestBody",
		Other: "Request body must not be empty",
	}

	MustBeOneOf = &i18n.Message{
		ID:    "MustBeOneOf",
		Other: "Must be one of: {{.Values}}",
	}

	GormDuplicateKey = &i18n.Message{
		ID:    "GormDuplicateKey",
		Other: "Value already exists",
	}
)
