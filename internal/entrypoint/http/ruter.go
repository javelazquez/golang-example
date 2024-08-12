package http

import (
	"golang-example/internal/entrypoint/http/handler"
	"golang-example/pkg/server"
	"golang-example/pkg/wrapper_result"
)

type Router struct {
	payoutHandler handler.PayoutHTTPHandlers
}

func NewRouter(payoutHandler handler.PayoutHTTPHandlers) *Router {
	return &Router{payoutHandler: payoutHandler}
}

func (r *Router) RouterURLs(app server.Web) {
	if r.payoutHandler != nil {
		payoutGroupV1 := app.Group("/v1")
		{
			payoutGroupV1.POST("/payout", wrapper_result.Wrapper(r.payoutHandler.Post))
			payoutGroupV1.GET("/payout/:id", wrapper_result.Wrapper(r.payoutHandler.Get))
		}

	}
}
