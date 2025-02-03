package main

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
)

func CreateLogger() *slog.Logger {
	// What does this do?
	ctx := context.Background()

	logExporter, err := otlploghttp.New(ctx)
	if err != nil {
		panic(err)
	}

	logProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	// makes sure resources are cleaned up when logProvider is not needed anymore
	defer logProvider.Shutdown(ctx)

	global.SetLoggerProvider(logProvider)

	logger := otelslog.NewLogger("hello",
		otelslog.WithSource(true),
	)

	return logger
}
