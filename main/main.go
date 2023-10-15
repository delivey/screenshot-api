package main

import (
	"log"
	"os"
	"screenshotapi/screenshot"

	"github.com/gofiber/fiber/v2"
)

type Url struct {
	Url string `json:"url"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/screenshot", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		body := new(Url)
		if err := c.BodyParser(body); err != nil {
			return err
		}

		url := body.Url
		hash := screenshot.GenerateHash(url)
		path := "tmp/" + hash + ".png"
		screenshot.MakeScreenshot(url, path)

		img, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		go screenshot.CleanupFromPath(path)

		c.Set("Content-Type", "image/png")
		return c.Send(img)
	})

	log.Fatal(app.Listen(":8080"))
}
