package errors

import "fmt"

// AppError standardizes error responses across services.
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Wrap constructs an AppError while preserving the underlying error.
func Wrap(code int, message string, err error) AppError {
	return AppError{Code: code, Message: message, Err: err}
}
