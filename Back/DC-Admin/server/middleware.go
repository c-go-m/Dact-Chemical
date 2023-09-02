package server

import (
	"github.com/c-go-m/DC-Admin/common/utility/constant"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func validateAccess(c *fiber.Ctx) error {
	return c.Next()
}

func controlError(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		switch useError.GetTypeError(err) {
		case constant.TypeGeneralError:
			c.Response().SetStatusCode(fiber.ErrInternalServerError.Code)
		case constant.TypeBussinesError:
			c.Response().SetStatusCode(fiber.ErrBadRequest.Code)
		default:
			c.Response().SetStatusCode(fiber.ErrInternalServerError.Code)
		}
	} else {
		c.Response().SetStatusCode(fiber.StatusOK)
	}
	return err
}

func getConfigCors() cors.Config {
	var configDefault = cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}

	return configDefault
}
