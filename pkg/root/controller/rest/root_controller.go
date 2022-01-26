package rest

import (
	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type RootController interface {
	Root(c *fiber.Ctx) error
}

type rootRestController struct{}

func NewRootRestController() RootController {
	return &rootRestController{}
}

func (r rootRestController) Root(c *fiber.Ctx) error {
	s := "Welcome To " + viper.GetString("name") + " : " + viper.GetString("version") + " : " + viper.GetString("profile")
	return response.Success(c, s)
}
