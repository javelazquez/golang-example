package appErrors

type InvalidParamError struct {
	baseError
}

func NewInvalidParamError(msg string, cause error) ApiError {
	return &InvalidParamError{baseError{
		code:          invalidParamError,
		message:       msg,
		originalError: cause,
	}}
}

type IdError struct {
	baseError
}

func NewIdError(msg string, cause error) ApiError {
	return &IdError{baseError{
		code:          invalidParamIDError,
		message:       msg,
		originalError: cause,
	}}
}
