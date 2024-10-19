package main

import (
	"context"
	_ "embed"
	"encoding/base64"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cretz/bine/tor"
	"github.com/cretz/bine/torutil/ed25519"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed assets/index.html
var html []byte

//go:embed assets/index.ans
var ansi []byte

var hiddenservice string

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

		if c.Hostname() != hiddenservice+".onion" {
			c.Set("onion-location", "http://"+hiddenservice+".onion")
		}

		if c.Get("X-Forwarded-Proto") == "http" {
			return c.Redirect(strings.Replace(c.BaseURL(), "http", "https", 1))
		}

		c.Set("Content-Type", "text/html")
		return c.Send(html)
	})

	app.Static("/", "assets", fiber.Static{Compress: true, MaxAge: 3600})

	go func() { log.Fatal(app.Listen(":4664")) }()
	//log.Fatal(app.Listen(":4664"))
	log.Println("Starting and registering onion service, please wait a couple of minutes...")
	t, err := tor.Start(context.Background(), nil) // &tor.StartConf{DebugWriter: os.Stdout}
	if err != nil {
		log.Panicf("Unable to start Tor: %v", err)
	}
	defer t.Close()

	listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer listenCancel()

	_key, _ := base64.StdEncoding.DecodeString(os.Getenv("HS_KEY"))
	onion, err := t.Listen(listenCtx, &tor.ListenConf{RemotePorts: []int{80}, Version3: true, Key: ed25519.PrivateKey(_key)})
	if err != nil {
		log.Panicf("Unable to create onion service: %v", err)
	}
	defer onion.Close()

	log.Printf("listening on http://%v.onion\n", onion.ID)
	hiddenservice = onion.ID

	log.Fatal(app.Listener(onion))
}
