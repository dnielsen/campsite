package tracing

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/openzipkin/zipkin-go"
	reporterHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

// TRACE_RECORD_RATE is a float from 0 to 1.
// 1 means 100% of traces will be recorded.
const TRACE_RECORD_RATE = 1

func NewTracer(serviceName string, servicePort string, c *config.Config) *zipkin.Tracer {
	endpointUrl := fmt.Sprintf("http://%v:9411/api/v2/spans", c.Tracing.Host)
	// The reporter sends traces to the zipkin server.
	reporter := reporterHttp.NewReporter(endpointUrl)

	// Local endpoint represents the local service information.
	localEndpoint, err := zipkin.NewEndpoint(serviceName, servicePort)
	if err != nil {
		log.Printf("Failed to create endpoint: %v", err)
	}

	// TRACE_RECORD_RATE is a float from 0 to 1.
	// 1 means 100% of traces will be recorded.
	sampler, err := zipkin.NewCountingSampler(TRACE_RECORD_RATE)
	if err != nil {
		log.Printf("Failed to create counting sampler: %v", err)
	}

	// Create a new tracer.
	t, err := zipkin.NewTracer(reporter, zipkin.WithSampler(sampler), zipkin.WithLocalEndpoint(localEndpoint))
	if err != nil {
		log.Fatalf("Failed to create tracer: %v", err)
		return nil
	}

	return t
}