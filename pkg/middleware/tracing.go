package middleware

import (
	"campsite/pkg/config"
	"campsite/pkg/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)


func Tracing(serviceName string, servicePort int, c *config.Config) mux.MiddlewareFunc {
	t := tracing.NewTracer(serviceName, servicePort, c)
	return zipkinHttp.NewServerMiddleware(t)
}