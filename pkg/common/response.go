package common

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Data    interface{} `json:"result"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SendJsonResponse(ctx *fiber.Ctx, httpCode int, message string, data interface{}) error {
	respPayload := BaseResponse{
		Data:    data,
		Message: message,
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(httpCode).JSON(respPayload)
}

func SendErrorResponse(ctx *fiber.Ctx, httpCode int, errMessage string) error {
	errResp := BaseResponse{
		Error: errMessage,
	}

	return ctx.Status(httpCode).JSON(errResp)
}
