package errutil

type errorItem struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type CustomError struct {
	Errors []errorItem `json:"errors"`
}

func (e *CustomError) Error() string {
	return e.Errors[0].Messages[0]
}

func buildErrorItem(field string, messages []string) errorItem {
	if field == "" {
		field = "detail"
	}
	return errorItem{
		Field:    field,
		Messages: messages,
	}
}

func New(field string, messages []string) *CustomError {
	error := buildErrorItem(field, messages)
	return &CustomError{
		Errors: []errorItem{error},
	}
}

func (e *CustomError) Add(field string, messages []string) *CustomError {
	error := buildErrorItem(field, messages)
	e.Errors = append(e.Errors, error)
	return e
}
