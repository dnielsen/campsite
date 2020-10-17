package middleware

import (
	"campsite/services/user/internal/config"
	"campsite/services/user/internal/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)


func Tracing(c *config.ServerConfig) mux.MiddlewareFunc {
	t := tracing.NewTracer(c)
	return zipkinHttp.NewServerMiddleware(t)
}