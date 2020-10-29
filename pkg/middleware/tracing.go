package middleware

import (
	"github.com/gorilla/mux"
	"github.com/openzipkin/zipkin-go"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
)

func Tracing(t *zipkin.Tracer) mux.MiddlewareFunc {
	return zipkinHttp.NewServerMiddleware(t)
}