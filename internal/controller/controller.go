package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/ports/service"
)

type controller struct {
	router      *fiber.App
	validate    *validator.Validate
	userService service.UserServices
}

type Controller interface {
	Init()
	SignUp(ctx *fiber.Ctx) error
}

func NewController(router *fiber.App, userService service.UserServices) Controller {
	validate := validator.New()
	return &controller{
		router, validate, userService,
	}
}

func (c *controller) Init() {
	api := c.router.Group("/api/v1")
	api.Post("/signup", c.SignUp)
	api.Get("/users/:uniqueId", c.FindUserByUniqueID)
}
