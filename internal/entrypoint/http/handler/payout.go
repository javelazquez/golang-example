package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		return resp.RequestResult{
			Status: http.StatusBadRequest,
			Body:   "rompio al parsear",
		}
	}

	err := p.payoutService.Create(c, requestBody)
	if err != nil {
		log.Error(fmt.Sprintf("fail to create payout: %s", err))
		return resp.RequestResult{
			Status: http.StatusInternalServerError,
			Body:   "rompio al guardar",
		}
	}

	return resp.RequestResult{
		Status: http.StatusCreated,
		Body:   nil,
	}
}

func (p *payoutHTTPHandlers) Get(c *gin.Context) resp.RequestResult {
	//TODO implement me
	panic("implement me")
}
