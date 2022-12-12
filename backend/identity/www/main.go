package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/google/wire"

	"github.com/lutanist/ste/backend/identity/core"
	"github.com/lutanist/ste/backend/identity/handlers"
)

func newApp(h *handlers.Handler) *fiber.App {
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

	h.Register(app)
	return app
}

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

	wire.Build(
		wire.Bind(new(core.AuthenticationScheme), new(*core.DefaultAuthenticationSchemeProvider)),
		core.NewDefaultAuthenticationSchemeProvider,
		core.NewSignInManager,
		handlers.NewHandler,
		newApp,
	)

	// TODO: initialize sign in manager

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
