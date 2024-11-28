package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(earlydata.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://maid.zone")
	})

	log.Fatal(app.Listen(":4664"))
}
