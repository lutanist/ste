// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initializer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/lutanist/ste/backend/identity/core"
	"github.com/lutanist/ste/backend/identity/handlers"
	"net/http"
)

// Injectors from wire.go:

func New() (*fiber.App, error) {
	defaultAuthenticationSchemeProvider := core.NewDefaultAuthenticationSchemeProvider()
	signInManager := core.NewSignInManager(defaultAuthenticationSchemeProvider)
	handler := handlers.NewHandler(signInManager)
	app := newApp(handler)
	return app, nil
}

// wire.go:

func newApp(h *handlers.Handler) *fiber.App {
	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Reload(true)

	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(recover2.New())

	h.Register(app)
	return app
}
