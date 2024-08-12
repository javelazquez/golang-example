package domain

type PayoutError struct {
	code       int    `json:"code"`
	message    string `json:"message"`
	statusCode int    `json:"status_code"`
}

func NewPayoutError(code int, message string, statusCode int) PayoutError {
	return PayoutError{code: code, message: message, statusCode: statusCode}
}

func (e PayoutError) GetMessage() string {
	return e.message
}

func (e PayoutError) GetStatusCode() int {
	return e.statusCode
}

func (e PayoutError) GetCode() int {
	return e.code
}
