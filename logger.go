package main

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func main() {
	// What does this do?
	ctx := context.Background()

	res, err := newResource()

	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	logProvider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	// makes sure resources are cleaned up when logProvider is not needed anymore
	defer logProvider.Shutdown(ctx)

	logger := otelslog.NewLogger("hello",
		otelslog.WithSource(true),
		otelslog.WithLoggerProvider(logProvider),
	)

	logger.Debug("Are you alive?")

	logger.Info("This is an info log message.")
	logger.Error("This is an error log message.")

	// Optionally, you can also log with additional context
	logger.Info("Log with context", slog.String("context.key", "context value"))

	// return logger
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName("my-service"),
			semconv.ServiceVersion("0.1.0"),
		))
}
