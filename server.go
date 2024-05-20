package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

type Article struct {
	Title string
	Text  string
}

func main() {

	sensitiveData := "some secret text"
	articles := []Article{}

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		slog.Info(sensitiveData)
		return c.SendString("Some other text")
	})

	app.Get("/articles", func(c fiber.Ctx) error {
		slog.Info(sensitiveData)
		return c.JSON(articles)
	})

	app.Listen(":3000")
	slog.Info("Server started on port 3000")
}
