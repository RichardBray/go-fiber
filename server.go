package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		slog.Info("Hello, World!")
		return c.SendString("Some other text")
	})

	app.Listen(":3000")
	slog.Info("Server started on port 3000")
}
