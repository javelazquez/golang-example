package appErrors

import (
	"errors"
	"golang-example/internal/core/domain"
	"net/http"
)

func HandlerError(err error) domain.PayoutError {
	var apiError ApiError
	statusCode := http.StatusInternalServerError

	switch {
	case errors.As(err, &apiError):
		switch apiError.(type) {
		case *InvalidParamError, *IdError, *InvalidRequestError, *AlreadyExistError,
			*ExternalIDError, *CountryError, *MerchantIDError, *CurrencyError, *AmountError:
			statusCode = http.StatusBadRequest
		case *NotFoundError:
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
	default:
		apiError = NewInternalError(err.Error(), err)
	}

	return domain.NewPayoutError(apiError.Code(), apiError.Error(), statusCode)
}
