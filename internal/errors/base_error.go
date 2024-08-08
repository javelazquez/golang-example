package errors

type baseError struct {
	code          int16  `json:"code"`
	message       string `json:"message"`
	originalError error  `json:"original_error"`
}

func (b baseError) Error() string {
	return b.message
}
