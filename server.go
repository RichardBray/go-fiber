package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("Hello, World!")
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
