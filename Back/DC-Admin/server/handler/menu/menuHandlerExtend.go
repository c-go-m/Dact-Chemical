package menu

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/gofiber/fiber/v2"
)

var urlBase = "/v1/api/menu/"

func (base *MenuHandler) AddHandlersGet() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "FindById/:Id": base.FindById,
		urlBase + "FindAll":      base.FindAll,
	}

	return handlers
}

func (base *MenuHandler) AddHandlersPost() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Create": base.Create,
	}

	return handlers
}

func (base *MenuHandler) AddHandlersPut() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Update": base.Update,
	}

	return handlers
}

func (base *MenuHandler) AddHandlersDelete() map[string]func(*fiber.Ctx) error {
	handlers := map[string]func(*fiber.Ctx) error{
		urlBase + "Delete/:Id": base.Delete,
	}

	return handlers
}

func (base *MenuHandler) parceObject(c *fiber.Ctx) (object entity.Menu, err error) {

	err = c.BodyParser(&object)

	if err != nil {
		return object, useError.ErrDataInvalid
	}
	return object, nil
}
