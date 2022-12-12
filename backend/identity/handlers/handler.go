package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lutanist/ste/backend/identity/core"
)

type Handler struct {
	sm *core.SignInManager
}

func NewHandler(sm *core.SignInManager) *Handler {
	return &Handler{
		sm: sm,
	}
}

func (h *Handler) Register(r *fiber.App) {
	ah := r.Group("/account")
	ah.Get("register", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title":     "Hello, World!",
			"ReturnUrl": "",
		}, "layouts/main")
	})

	ah.Get("login", func(c *fiber.Ctx) error {
		returnUrl := c.Query("returnUrl", "~/")

		return c.Render("login", fiber.Map{
			"ExternalLogins": h.sm.GetExternalAuthenticationSchemes(),

			"Title":     "Hello, World!",
			"ReturnUrl": returnUrl,
		}, "layouts/main")
	})
}
