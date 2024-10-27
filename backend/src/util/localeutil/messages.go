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

	InvalidState = &i18n.Message{
		ID:    "InvalidState",
		Other: "Invalid state",
	}

	AuthorizationCodeNotFound = &i18n.Message{
		ID:    "AuthorizationCodeNotFound",
		Other: "Authorization code not found",
	}

	CannotExchangeAuthorizationCode = &i18n.Message{
		ID:    "CannotExchangeAuthorizationCode",
		Other: "Can not exchange authorization code for tokens",
	}

	FailedToFetchJWKS = &i18n.Message{
		ID:    "FailedToFetchJWKS",
		Other: "Failed to fetch JWKS",
	}

	FailedToParseToken = &i18n.Message{
		ID:    "FailedToParseToken",
		Other: "Failed to parse token",
	}

	NoKidFieldInJWTTokenHeader = &i18n.Message{
		ID:    "NoKidFieldInJWTTokenHeader",
		Other: "No 'kid' field in JWT token header",
	}

	UnableToFindKeyWithKid = &i18n.Message{
		ID:    "UnableToFindKeyWithKid",
		Other: "Unable to find key with kid",
	}

	FailedToCreateRawKey = &i18n.Message{
		ID:    "FailedToCreateRawKey",
		Other: "Failed to create raw key",
	}

	ExpectedRSAKey = &i18n.Message{
		ID:    "ExpectedRSAKey",
		Other: "Expected RSA public key but got something else",
	}

	FailedToVerifyToken = &i18n.Message{
		ID:    "FailedToVerifyToken",
		Other: "Failed to verify token",
	}

	TokenHasExpired = &i18n.Message{
		ID:    "TokenHasExpired",
		Other: "Token has expired",
	}

	NoRealmFound = &i18n.Message{
		ID:    "NoRealmFound",
		Other: "No realm found",
	}

	RefreshTokenNotFound = &i18n.Message{
		ID:    "RefreshTokenNotFound",
		Other: "Refresh token not found",
	}

	CannotExchangeRefreshToken = &i18n.Message{
		ID:    "CannotExchangeRefreshToken",
		Other: "Can not exchange refresh token for tokens",
	}

	Unauthorized = &i18n.Message{
		ID:    "Unauthorized",
		Other: "Unauthorized",
	}
)
