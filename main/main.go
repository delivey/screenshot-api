package mainn

import (
	"log"
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
		screenshot.MakeScreenshot(url)

		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
