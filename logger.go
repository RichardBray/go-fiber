package main

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func CreateLogger(ctx context.Context) (*slog.Logger, *log.LoggerProvider) {
	res, err := newResource()
	if err != nil {
		fmt.Errorf("failed to create resource: %w", err)
	}

	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	logProvider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	logger := otelslog.NewLogger("hello",
		otelslog.WithSource(true), // Includes source location of log in attributes
		otelslog.WithLoggerProvider(logProvider),
	)

	return logger, logProvider
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName("go-fiber-app"),
			semconv.ServiceVersion("0.1.0"),
		))
}
