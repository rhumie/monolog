package exceptions

import "fmt"

// AuthenticationException is the exception that occurs when authentication is invalid.
type AuthenticationException struct {
	Code    int
	Message string
}

func (err *AuthenticationException) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}
