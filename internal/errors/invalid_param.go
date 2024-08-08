package errors

type InvalidParamError struct {
	baseError
}

func NewInvalidParamError(msg string, cause error) error {
	return &InvalidParamError{baseError{
		code:          invalidParamError,
		message:       msg,
		originalError: cause,
	}}
}

func (g InvalidParamError) Error() string {
	return g.baseError.message
}
