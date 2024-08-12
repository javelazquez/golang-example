package appErrors

type AlreadyExistError struct {
	baseError
}

func NewAlreadyExistError(msg string) ApiError {
	return &AlreadyExistError{baseError{
		code:    payoutAlreadyExist,
		message: msg,
	}}
}
