package appErrors

const (
	genericErrorCode int = 1000

	payoutNotFound     int = 1001
	payoutAlreadyExist int = 1002

	invalidParamError   int = 2000
	invalidParamIDError int = 2000

	invalidRequest           int = 3000
	invalidRequestExternalID int = 3001
	invalidRequestMerchantID int = 3002
	invalidRequestAmount     int = 3003
	invalidRequestCountry    int = 3004
	invalidRequestCurrency   int = 3005

	internalErrorCode int = 5000
)
