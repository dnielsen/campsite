package middleware

import (
	"campsite/services/session/internal/config"
	"campsite/services/session/internal/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)

func Tracing(c *config.ServerConfig) mux.MiddlewareFunc {
	t := tracing.NewTracer(c)
	return zipkinHttp.NewServerMiddleware(t)
}
