package wrapper_result

import (
	"github.com/gin-gonic/gin"
)

// RequestResult represents the result of a router request
type RequestResult struct {
	// Status is the htt status code of the response
	Status int
	// Body holds the response body of the request
	Body any
}

// HandleFunc is a function to handle the main logic and returns the results of the request
type HandleFunc func(*gin.Context) RequestResult

// Wrapper is a function to handle the response of the requests
func Wrapper(f HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := f(c)
		c.JSON(r.Status, r.Body)
	}
}
