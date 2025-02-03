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

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hallo Welt!!")
	})

	app.Get("/articles", func(c fiber.Ctx) error {
		logger.Info("Getting all articles")
		return c.JSON(articles)
	})

	app.Post("/article", func(c fiber.Ctx) error {
		article := new(Article)
		if err := c.Bind().Body(article); err != nil {
			logger.Error("Error parsing article", err)
			return err
		}
		articles = append(articles, *article)
		logger.Info("Added new article")
		return c.JSON(article)
	})

	PORT := ":3000"
	app.Listen(PORT)
}
