package middleware

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/tracing"
	"github.com/gorilla/mux"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)


func Tracing(serviceName string, servicePort int, c *config.Config) mux.MiddlewareFunc {
	t := tracing.NewTracer(serviceName, string(rune(servicePort)), c)
	return zipkinHttp.NewServerMiddleware(t)
}