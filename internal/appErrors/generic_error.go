package appErrors

type GenericError struct {
	baseError
}

func NewGenericError(msg string, cause error) ApiError {
	return &GenericError{baseError{
		code:          genericErrorCode,
		message:       msg,
		originalError: cause,
	}}
}
