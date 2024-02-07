package trace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

const traceEndPoint string = "http://127.0.0.1:14268/api/traces"

func Init(env, serviceName string) {
	err := SetTracerProvider(env, serviceName)
	if err != nil {
		panic(err)
	}
}

func SetTracerProvider(env, serviceName string) error {
	// create the jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(traceEndPoint)))
	if err != nil {
		return err
	}

	// New trace provider
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// always be sure to batch in production
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName), // service name,实例名称
				attribute.String("env", env),               // environment
				//attribute.String("ID", serviceName), // version
			)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
