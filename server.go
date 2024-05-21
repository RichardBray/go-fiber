package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

type Article struct {
	Id    int
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
		return c.SendString("Nothing to see here...")
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

	app.Delete("/article/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		for i, article := range articles {
			if id == fmt.Sprint(article.Id) {
				articles = append(articles[:i], articles[i+1:]...)
				return c.SendString("Article deleted")
			}
		}
		return c.SendString("Article not found")
	})

	PORT := "4000"
	app.Listen(PORT)
	slog.Info("Server started on port:" + PORT)
}
