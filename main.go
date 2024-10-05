package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed assets/index.html
var html []byte

//go:embed assets/index.ans
var ansi []byte

func main() {
	app := fiber.New()
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(earlydata.New())

	app.Get("/", func(c *fiber.Ctx) error {
		ua := strings.ToLower(c.Get("User-Agent"))
		if len(ua) > 5 {
			switch ua[:4] {
			case "curl", "wget":
				return c.Send(ansi)
			}
		}

		if c.Get("X-Forwarded-Proto") == "http" {
			return c.Redirect(strings.Replace(c.BaseURL(), "http", "https", 1))
		}

		c.Set("Content-Type", "text/html")
		return c.Send(html)
	})

	app.Static("/", "assets", fiber.Static{Compress: true, MaxAge: 3600})

	log.Fatal(app.Listen(":4664"))
}
