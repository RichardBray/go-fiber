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
	articles := []Article{}

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Some other text")
	})

	app.Get("/article", func(c fiber.Ctx) error {
		return c.JSON(articles)
	})

	app.Post("/article", func(c fiber.Ctx) error {
		article := new(Article)
		if err := c.Bind().Body(article); err != nil {
			return err
		}
		articles = append(articles, *article)
		return c.JSON(article)
	})

	app.Listen(":3000")
	slog.Info("Server started on port 3000")
}
