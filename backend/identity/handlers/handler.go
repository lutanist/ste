package handlers

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func (h *Handler) Register(r *fiber.App) {
	ah := r.Group("/account")
	ah.Get("register", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title":     "Hello, World!",
			"ReturnUrl": "",
		}, "layouts/main")
	})

	ah.Get("login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"Title":     "Hello, World!",
			"ReturnUrl": "",
		}, "layouts/main")
	})
}
