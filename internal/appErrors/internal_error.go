package appErrors

type InternalError struct {
	baseError
}

func NewInternalError(msg string, cause error) ApiError {
	return &InternalError{baseError{
		code:          internalErrorCode,
		message:       msg,
		originalError: cause,
	}}
}
