package auth

type ApiKeyError struct {
	Message string
}

func (e *ApiKeyError) Error() string {
	return e.Message
}

func NewApiKeyError(message string) error {
	return &ApiKeyError{Message: message}
}
