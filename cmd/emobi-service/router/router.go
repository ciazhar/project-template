package router

import (
	"github.com/ciazhar/emobi-service/cmd/emobi-service/app"
	"github.com/ciazhar/emobi-service/pkg/file"
	"github.com/ciazhar/emobi-service/pkg/root"
	"github.com/ciazhar/emobi-service/pkg/vehicle"

	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Init(a app.Application) error {

	// Default error handler
	var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
		// Default unknown error
		code := response.CodeInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			// Override tracksolid code if fiber.Error type
			code = e.Code
		} else {
			// log error
			sentry.CaptureException(err)
		}

		// Set Content-Type: text/plain; charset=utf-8
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		// Return statuscode with error message
		return c.Status(code).JSON(response.Response{
			Message: err.Error(),
		})
	}

	//init fiber and middleware
	r := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})
	r.Use(cors.New(cors.Config{AllowCredentials: true}))
	r.Use(pprof.New())
	r.Use(recover.New())
	a.Router = r

	root.Init(a)
	vehicle.Init(a)
	file.Init(a)

	r.Static("/download", viper.GetString("path"))

	//route 404
	r.Use(func(c *fiber.Ctx) error {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "route not found",
		})
	})

	//run
	if viper.GetString("profile") != "debug" {
		sentry.CaptureMessage("appplication start in port : " + viper.GetString("port"))
	}
	return r.Listen(":" + viper.GetString("port"))
}
