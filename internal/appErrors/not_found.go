package appErrors

type NotFoundError struct {
	baseError
}

func NewNotFoundError(msg string) ApiError {
	return &NotFoundError{baseError{
		code:    payoutNotFound,
		message: msg,
	}}
}
