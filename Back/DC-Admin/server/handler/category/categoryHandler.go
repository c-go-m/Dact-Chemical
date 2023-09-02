package category

import (
	"strconv"

	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/c-go-m/DC-Admin/service/category"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service category.ICategoryService
}

func New(service category.ICategoryService) *CategoryHandler {
	handler := CategoryHandler{
		service: service,
	}
	return &handler
}

func (base *CategoryHandler) FindAll(c *fiber.Ctx) error {
	result, err := base.service.FindAll()

	if err != nil {
		return useError.ErrorInitServer
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (base *CategoryHandler) Create(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	err = base.service.Create(&object)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	return c.Status(fiber.StatusOK).JSON(true)
}

func (base *CategoryHandler) Update(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	return c.Status(fiber.StatusOK).JSON(base.service.Update(&object))
}

func (base *CategoryHandler) Delete(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	err = base.service.Delete(&object)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	return c.Status(fiber.StatusOK).JSON(true)
}

func (base *CategoryHandler) FindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("Id"))
	result, err := base.service.FindById(id)

	if err != nil {
		return useError.ErrorInitServer
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
