package http

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
)

func CreateHttpServer(ctx context.Context) {
	app := fiber.New()

	log.Println("Starting http server")

	app.Get("/", helloWorld)

	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) {
	c.SendString("Hello world")
}
