//go:build wireinject
// +build wireinject

package initializer

import (
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

func New() (*fiber.App, error) {
	wire.Build(
		core.NewDefaultAuthenticationSchemeProvider,
		wire.Bind(new(core.AuthenticationSchemeProvider), new(*core.DefaultAuthenticationSchemeProvider)),
		core.NewSignInManager,
		handlers.NewHandler,
		newApp,
	)

	return nil, nil
}
