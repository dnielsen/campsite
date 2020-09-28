package tracing

import (
	"dave-web-app/packages/event-service/internal/config"
	"github.com/gorilla/mux"
	"github.com/openzipkin/zipkin-go"
	zipkinHttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/openzipkin/zipkin-go/model"
	reporterHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
	"strconv"
)

const ENDPOINT_URL = "http://localhost:9411/api/v2/spans"
const SERVICE_NAME = "event-service"
// TRACE_RECORD_RATE is a float from 0 to 1.
// 1 means 100% of traces will be recorded.
const TRACE_RECORD_RATE = 1

func NewTracer(c *config.ServerConfig) *zipkin.Tracer {
	// The reporter sends traces to the zipkin server.
	reporter := reporterHttp.NewReporter(ENDPOINT_URL)

	// Convert the port to an integer.
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		log.Fatalf("Failed to convert port to integer: %v", err)
	}

	// Local endpoint represents the local service information.
	// The port needs to be of type `uint16` so we're converting it.
	localEndpoint := &model.Endpoint{ServiceName: SERVICE_NAME, Port: uint16(port)}

	// TRACE_RECORD_RATE is a float from 0 to 1.
	// 1 means 100% of traces will be recorded.
	sampler, err := zipkin.NewCountingSampler(TRACE_RECORD_RATE)
	if err != nil {
		log.Fatalf("Failed to create counting sampler: %v", err)
	}

	// Create a new tracer.
	t, err := zipkin.NewTracer(reporter, zipkin.WithSampler(sampler), zipkin.WithLocalEndpoint(localEndpoint))
	if err != nil {
		log.Fatalf("Failed to create tracer: %v", err)
	}

	return t
}

func NewTracingMiddleware(t *zipkin.Tracer) mux.MiddlewareFunc {
	return zipkinHttp.NewServerMiddleware(t)
}