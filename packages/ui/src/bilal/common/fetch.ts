import { Tracer, BatchRecorder, jsonEncoder } from "zipkin";
const CLSContext = require("zipkin-context-cls");
const { HttpLogger } = require("zipkin-transport-http");
const wrapFetch = require("zipkin-instrumentation-fetch");

// Setup the tracer
const tracer = new Tracer({
  ctxImpl: new CLSContext("zipkin"), // implicit in-process context
  recorder: new BatchRecorder({
    logger: new HttpLogger({
      endpoint: "http://localhost:9411/api/v2/spans",
      jsonEncoder: jsonEncoder.JSON_V2,
    }),
  }), // batched http recorder
  localServiceName: "ui", // name of this application
});

export const zipkinFetch = wrapFetch(fetch, { tracer });
