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

// var alphanum = [62]rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'Z', 'X', 'C', 'V', 'B', 'N', 'M', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// func RandomSlug(length int) string {
// 	b := make([]rune, length)
// 	for i := range b {
// 		b[i] = alphanum[rand.Intn(len(alphanum))]
// 	}

// 	return string(b)
// }

// var isSavingShortened bool

// func saveShortened() {
// 	if isSavingShortened {
// 		return
// 	}
// 	isSavingShortened = true

// 	raw, err := json.Marshal(shortened)
// 	if err != nil {
// 		isSavingShortened = false
// 		panic(err)
// 	}

// 	err = os.WriteFile(data+"shortened.json", raw, 0666)
// 	isSavingShortened = false
// 	if err != nil {
// 		panic(err)
// 	}
// }

//go:embed assets/index.html
var html []byte

//go:embed assets/index.ans
var ansi []byte

// type Shortened struct {
// 	SlugToLink       mut.Map[string, string]
// 	LinkToRandomSlug mut.Map[string, string]
// }

//	var shortened = Shortened{
//		SlugToLink: mut.Map[string, string]{
//			Map: make(map[string]string),
//			Mut: &sync.RWMutex{},
//		},
//		LinkToRandomSlug: mut.Map[string, string]{
//			Map: make(map[string]string),
//			Mut: &sync.RWMutex{},
//		},
//	}
//
// var slugre = regexp.MustCompile(`^[a-zA-Z0-9]{1,256}$`)
// var json = jsoniter.ConfigFastest

// const data = "/data/"

func main() {
	// go func() {
	// 	raw, err := os.ReadFile(data + "shortened.json")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	err = json.Unmarshal(raw, &shortened)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
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

	app.Get("/s", func(c *fiber.Ctx) error {
		ua := strings.ToLower(c.Get("User-Agent"))
		if len(ua) > 5 {
			switch ua[:4] {
			case "curl", "wget":
				return c.SendFile("assets/s.ans", true)
			}
		}

		if c.Get("X-Forwarded-Proto") == "http" {
			return c.Redirect(strings.Replace(c.BaseURL(), "http", "https", 1))
		}

		return c.SendFile("assets/s.html", true)
	})

	// app.Get("/s/:slug", func(c *fiber.Ctx) error {
	// 	slug := c.Params("slug")
	// 	val, ok := shortened.SlugToLink.GetValueAndState(slug)
	// 	if !ok {
	// 		return fiber.ErrNotFound
	// 	}

	// 	return c.Redirect(val)
	// })

	// app.Get("/api/shorten", limiter.New(limiter.Config{SkipFailedRequests: true, Expiration: time.Second, Max: 2, LimiterMiddleware: limiter.SlidingWindow{}}), func(c *fiber.Ctx) error {
	// 	s := c.Query("url")
	// 	if s == "" || len(s) > 256 {
	// 		return fiber.ErrBadRequest
	// 	}

	// 	if parsed, err := url.Parse(s); err != nil || parsed.Scheme == "" {
	// 		return fiber.ErrBadRequest
	// 	}

	// 	slug := strings.Trim(c.Query("slug"), " ")
	// 	if !slugre.MatchString(slug) && slug != "" {
	// 		return fiber.ErrBadRequest
	// 	}

	// 	if _, ok := shortened.SlugToLink.GetValueAndState(slug); ok {
	// 		return c.Status(fiber.StatusConflict).SendString("slug taken")
	// 	}

	// 	if slug == "" {
	// 		if slug, ok := shortened.LinkToRandomSlug.GetValueAndState(s); ok {
	// 			return c.SendString(slug)
	// 		}

	// 		for i := 0; i < 6; i++ {
	// 			if i == 5 {
	// 				return fiber.ErrInternalServerError // fuck it, we're out of luck this time
	// 			}

	// 			slug = RandomSlug(6)
	// 			if _, ok := shortened.SlugToLink.GetValueAndState(slug); !ok {
	// 				cls := strings.Clone(slug)
	// 				clu := strings.Clone(s)

	// 				shortened.LinkToRandomSlug.Set(clu, cls)
	// 				shortened.SlugToLink.Set(cls, clu)
	// 				log.Printf("shortened %s -> %s\n", s, slug)
	// 				return c.SendString(slug)
	// 			}
	// 		}
	// 	}

	// 	shortened.SlugToLink.Set(strings.Clone(slug), strings.Clone(s))
	// 	go saveShortened()
	// 	log.Printf("shortened %s -> %s\n", s, slug)
	// 	return c.SendString(slug)
	// })

	app.Static("/", "assets", fiber.Static{Compress: true, MaxAge: 3600})

	log.Fatal(app.Listen(":4664"))
}
