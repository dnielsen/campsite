package middleware

import (
	"campsite/packages/event-service/internal/config"
	"campsite/packages/event-service/internal/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// Tracing middleware enables tracing provided by the Hypertrace which
// provides services such as Zipkin.
func Tracing(c *config.ServerConfig) mux.MiddlewareFunc {
	t := tracing.NewTracer(c)
	return zipkinHttp.NewServerMiddleware(t)
}