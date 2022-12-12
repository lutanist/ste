//go:build wireinject
// +build wireinject

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lutanist/ste/backend/identity/initializer"
)

func main() {
	app, err := initializer.New()
	if err != nil {
		log.Fatal(err)
	}
	if app == nil {
		log.Fatal("xxx")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/account/register", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
