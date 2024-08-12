package appErrors

type InvalidRequestError struct {
	baseError
}

func NewInvalidRequestError(msg string, cause error) ApiError {
	return &InvalidRequestError{baseError{
		code:          invalidRequest,
		message:       msg,
		originalError: cause,
	}}
}

type ExternalIDError struct {
	baseError
}

func NewExternalIDError(msg string, cause error) ApiError {
	return &ExternalIDError{baseError{
		code:          invalidRequestExternalID,
		message:       msg,
		originalError: cause,
	}}
}

type CountryError struct {
	baseError
}

func NewCountryError(msg string, cause error) ApiError {
	return &CountryError{baseError{
		code:          invalidRequestCountry,
		message:       msg,
		originalError: cause,
	}}
}

type MerchantIDError struct {
	baseError
}

func NewMerchantIDError(msg string, cause error) ApiError {
	return &MerchantIDError{baseError{
		code:          invalidRequestMerchantID,
		message:       msg,
		originalError: cause,
	}}
}

type AmountError struct {
	baseError
}

func NewAmountError(msg string, cause error) ApiError {
	return &AmountError{baseError{
		code:          invalidRequestAmount,
		message:       msg,
		originalError: cause,
	}}
}

type CurrencyError struct {
	baseError
}

func NewCurrencyError(msg string, cause error) ApiError {
	return &CurrencyError{baseError{
		code:          invalidRequestCurrency,
		message:       msg,
		originalError: cause,
	}}
}
