package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/dto/request"
	"github.com/wahyurudiyan/super-sharing/user-svc/pkg/common"
)

func (c *controller) SignUp(ctx *fiber.Ctx) error {
	var request request.SignUp
	err := ctx.BodyParser(&request)
	if err != nil {
		return common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err = c.validate.Struct(request)
	if err != nil {
		return common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if err := c.userService.SignUp(ctx.Context(), request); err != nil {
		return common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return common.SendJsonResponse(ctx, 200, "OK", "signup success")
}
