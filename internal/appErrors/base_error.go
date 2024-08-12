package appErrors

type ApiError interface {
	Error() string
	Code() int
	Unwrap() error
}

type baseError struct {
	code          int    `json:"code"`
	message       string `json:"message"`
	originalError error  `json:"original_error"`
}

func (b baseError) Error() string {
	return b.message
}

func (b baseError) Code() int {
	return b.code
}

func (b baseError) Unwrap() error {
	return b.originalError
}
