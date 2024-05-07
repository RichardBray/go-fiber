package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("Hello, World Again!")
		return c.SendString("Some other text")
	})

	app.Listen(":3000")
}
