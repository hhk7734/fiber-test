package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	zap.L().Fatal("failed to start server", zap.Error(app.Listen(":8080")))
}
