package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		log.Println("Hello, World!")
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
	log.Println("Server started on port 3000")
}
