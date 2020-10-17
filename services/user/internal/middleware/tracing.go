package middleware

import (
	"campsite/packages/user/internal/config"
	"campsite/packages/user/internal/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)


func Tracing(c *config.ServerConfig) mux.MiddlewareFunc {
	t := tracing.NewTracer(c)
	return zipkinHttp.NewServerMiddleware(t)
}