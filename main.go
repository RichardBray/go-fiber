package main

import (
	"github.com/gofiber/fiber/v3"
)

type Article struct {
	Title string
	Text  string
}

func setupApp() *fiber.App {
	articles := []Article{}
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hola World")
	})

	app.Get("/articles", func(c fiber.Ctx) error {
		return c.JSON(articles)
	})

	app.Post("/article", func(c fiber.Ctx) error {
		article := new(Article)
		if err := c.Bind().Body(article); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{"error": "Invalid input"},
			)
		}
		articles = append(articles, *article)
		return c.JSON(article)
	})

	return app
}

func main() {
	app := setupApp()
	PORT := ":3000"
	app.Listen(PORT)
}
