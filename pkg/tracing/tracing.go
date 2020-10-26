package tracing

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	reporterHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

// TRACE_RECORD_RATE is a float from 0 to 1.
// 1 means 100% of traces will be recorded.
const TRACE_RECORD_RATE = 1

func NewTracer(serviceName string, servicePort int, c *config.Config) *zipkin.Tracer {
	endpointUrl := fmt.Sprintf("http://%v:9411/api/v2/spans", c.Tracing.Host)
	// The reporter sends traces to the zipkin server.
	reporter := reporterHttp.NewReporter(endpointUrl)

	// Local endpoint represents the local service information.
	// The port needs to be of type `uint16` so we're converting it.
	localEndpoint := &model.Endpoint{
		ServiceName: serviceName,
		Port:        uint16(servicePort),
	}

	// TRACE_RECORD_RATE is a float from 0 to 1.
	// 1 means 100% of traces will be recorded.
	sampler, err := zipkin.NewCountingSampler(TRACE_RECORD_RATE)
	if err != nil {
		log.Fatalf("Failed to create counting sampler: %v", err)
	}

	// Create a new tracer.
	t, err := zipkin.NewTracer(reporter, zipkin.WithSampler(sampler), zipkin.WithLocalEndpoint(localEndpoint), zipkin.WithTags(map[string]string{
		"hi": "hello",
	}))
	if err != nil {
		log.Fatalf("Failed to create tracer: %v", err)
	}

	return t
}