package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"

	"github.com/lutanist/ste/backend/identity/handlers"
)

func main() {
	engine := html.NewFileSystem(http.Dir("./views"), ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true)
	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: false,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	h := handlers.Handler{}
	h.Register(app)

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
