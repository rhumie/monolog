package services

// ServiceError is s
type ServiceError interface {
	error
	Code() string
	Message() string
	UnAuthorized() bool
}

type baseError struct {
	code         string
	message      string
	unAuthorized bool
	err          error
}

func (e *baseError) Code() string {
	return e.code
}

func (e *baseError) Message() string {
	return e.message
}

func (e *baseError) UnAuthorized() bool {
	return e.unAuthorized
}

func (e *baseError) Error() string {
	return "[" + e.code + "]" + " " + e.message + ":" + e.err.Error()
}

func newAuthenticationError(err error) ServiceError {
	return &baseError{
		code:         "ServiceUnAuthorized",
		message:      "service authentication failed.",
		unAuthorized: true,
		err:          err,
	}
}
