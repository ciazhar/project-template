package response

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    Code        `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Code int

const (
	CodeSuccess             Code = 200
	CodeBadRequest               = 400
	CodeInternalServerError      = 500
	CodeUnauthorized             = 401
)

func Data(data interface{}) Response {
	return Response{
		Message: "succes",
		Code:    CodeSuccess,
		Data:    data,
	}
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(Data(data))
}

func Error(err error, code ...Code) error {
	if err == nil {
		return nil
	}

	if code != nil && len(code) == 1 {
		return fiber.NewError(int(code[0]), err.Error())
	} else {
		sentry.CaptureException(err)
		return fiber.NewError(CodeInternalServerError, err.Error())
	}
}
