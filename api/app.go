package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func StartWebApp() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	app.Get("/metrics", monitor.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"path": "v2",
		})
	})
	v2Router(app.Group("/v2"))

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
