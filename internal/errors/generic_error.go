package errors

type GenericError struct {
	baseError
}

func NewGenericError(msg string, cause error) error {
	return &GenericError{baseError{
		code:          genericErrorCode,
		message:       msg,
		originalError: cause,
	}}
}

func (g GenericError) Error() string {
	return g.baseError.message
}
