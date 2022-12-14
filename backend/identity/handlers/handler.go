package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lutanist/ste/backend/identity/core"
	"github.com/lutanist/ste/backend/identity/ent"
)

type RegisterInput struct {
	Email           string `json:"email,omitempty" form:"email"`
	Password        string `json:"password,omitempty" form:"password"`
	ConfirmPassword string `json:"confirm_password,omitempty" form:"confirm_password"`
	Name            string `json:"name,omitempty" form:"name"`
}

type Handler struct {
	client *ent.Client
	sm     *core.SignInManager
	um     *core.UserManager
}

func NewHandler(sm *core.SignInManager, um *core.UserManager) *Handler {
	return &Handler{
		sm: sm,
		um: um,
	}
}

func (h *Handler) Register(r *fiber.App) {
	ag := r.Group("/account")

	ag.Get("register", func(c *fiber.Ctx) error {
		input := new(RegisterInput)
		if err := c.BodyParser(input); err != nil {
			return err
		}

		newUser := ent.User{
			Name:     input.Name,
			Username: input.Email,
		}

		result, err := h.um.Create(&newUser, input.Password)
		if err != nil {
			return err
		}

		if result.Succeeded {

		}

		return c.Render("signup", fiber.Map{
			"ExternalLogins": h.sm.GetExternalAuthenticationSchemes(),
			"Title":          "Hello, World!",
			"ReturnUrl":      "",
		}, "layouts/main")
	})

	ag.Get("login", func(c *fiber.Ctx) error {
		returnUrl := c.Query("returnUrl", "~/")

		return c.Render("login", fiber.Map{
			"ExternalLogins": h.sm.GetExternalAuthenticationSchemes(),

			"Title":     "Hello, World!",
			"ReturnUrl": returnUrl,
		}, "layouts/main")
	})
}
