package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyurudiyan/super-sharing/user-svc/pkg/common"
)

func (c *controller) FindUserByUniqueID(ctx *fiber.Ctx) error {
	uniqueId := ctx.Params("uniqueId")
	if uniqueId == "" {
		return common.SendErrorResponse(ctx, http.StatusBadRequest, "unique id cannot be null")
	}

	user, err := c.userService.FindUserByUniqueID(ctx.Context(), uniqueId)
	if err != nil {
		return common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return common.SendJsonResponse(ctx, http.StatusOK, "OK", user)
}
