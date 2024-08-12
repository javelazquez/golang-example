package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-example/internal/appErrors"
	"golang-example/internal/core/ports"
	"golang-example/internal/entrypoint/model"
	"golang-example/pkg/logger"
	resp "golang-example/pkg/wrapper_result"
	"net/http"
)

type PayoutHTTPHandlers interface {
	Post(c *gin.Context) resp.RequestResult
	Get(c *gin.Context) resp.RequestResult
}

type payoutHTTPHandlers struct {
	payoutService ports.PayoutService
}

func NewPayoutHTTPHandlers(payoutService ports.PayoutService) *payoutHTTPHandlers {
	return &payoutHTTPHandlers{payoutService: payoutService}
}

func (p *payoutHTTPHandlers) Post(c *gin.Context) resp.RequestResult {
	log := logger.NewLog()
	var requestBody model.CreatePayoutRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Warn(fmt.Sprintf("fail to parse body: %s", err))
		pErr := appErrors.HandlerError(appErrors.NewInvalidRequestError(err.Error(), err))
		return resp.RequestResult{
			Status: pErr.GetStatusCode(),
			Body:   model.NewErrorResponse(pErr.GetCode(), pErr.GetMessage()),
		}
	}

	pID, err := p.payoutService.Create(c, requestBody)
	if err != nil {
		log.Error(fmt.Sprintf("fail to create payout: %s", err))
		pErr := appErrors.HandlerError(err)
		return resp.RequestResult{
			Status: pErr.GetStatusCode(),
			Body:   model.NewErrorResponse(pErr.GetCode(), pErr.GetMessage()),
		}
	}

	return resp.RequestResult{
		Status: http.StatusCreated,
		Body:   model.NewCreatePayoutResponse(pID),
	}
}

func (p *payoutHTTPHandlers) Get(c *gin.Context) resp.RequestResult {
	log := logger.NewLog()
	id := c.Param("id")
	requestParam := model.NewGetPayoutRequest(id)

	payoutResponse, err := p.payoutService.Get(c, requestParam)
	if err != nil {
		log.Error(fmt.Sprintf("fail to get payout: %s", err))
		pErr := appErrors.HandlerError(err)
		return resp.RequestResult{
			Status: pErr.GetStatusCode(),
			Body:   model.NewErrorResponse(pErr.GetCode(), pErr.GetMessage()),
		}
	}

	return resp.RequestResult{
		Status: http.StatusOK,
		Body:   model.NewGetPayoutResponse(payoutResponse),
	}
}
