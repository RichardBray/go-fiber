package main

import (
	"github.com/gofiber/fiber/v3"
)

type Article struct {
	Title string
	Text  string
}

func main() {
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
			return err
		}
		articles = append(articles, *article)
		return c.JSON(article)
	})

	PORT := ":3000"
	app.Listen(PORT)
}
