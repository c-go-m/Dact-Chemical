package product

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/gofiber/fiber/v2"
)

var urlBase = "/v1/api/product/"

func (base *ProductHandler) AddHandlersGet() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "FindById/:Id": base.FindById,
		urlBase + "FindAll":      base.FindAll,
	}

	return handlers
}

func (base *ProductHandler) AddHandlersPost() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Create": base.Create,
	}

	return handlers
}

func (base *ProductHandler) AddHandlersPut() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Update": base.Update,
	}

	return handlers
}

func (base *ProductHandler) AddHandlersDelete() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Delete/:Id": base.Delete,
	}

	return handlers
}

func (base *ProductHandler) parceObject(c *fiber.Ctx) (object entity.Product, err error) {

	err = c.BodyParser(&object)

	if err != nil {
		return object, useError.ErrDataInvalid
	}
	return object, nil
}
