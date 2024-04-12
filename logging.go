package main

import (
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	log.Println("Non structured log")

	// logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level:     slog.LevelDebug,
	// 	AddSource: true,
	// })

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().Format(time.ANSIC))
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.Group("values",
			slog.String("app", "example"),
			slog.Int("userId", 42),
		),
	})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)

	// Default log levels
	slog.Info("Structured log")
	slog.Debug("debug level")
	slog.Warn("warn level")
	slog.Error("error level")

	// Custom logger
	logger.Debug("debug level")
	// logs with attributes
	logger.Debug("debug level", "app", "example", "userId", 42)
	logger.Debug("debug level", slog.String("app", "example"))
	logger.Debug("debug level", slog.Group("values",
		slog.String("app", "example"),
		slog.Int("userId", 42),
	))
}
