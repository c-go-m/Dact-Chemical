package attached

import (
	"fmt"
	"strconv"

	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/c-go-m/DC-Admin/service/attached"
	"github.com/gofiber/fiber/v2"
)

type AttachedHandler struct {
	service attached.IAttachedService
}

func New(service attached.IAttachedService) *AttachedHandler {
	handler := AttachedHandler{
		service: service,
	}
	return &handler
}

func (base *AttachedHandler) FindAll(c *fiber.Ctx) error {
	fmt.Println("findAll 1")
	result, err := base.service.FindAll()

	if err != nil {
		return useError.ErrorInitServer
	}
	fmt.Println("findAll 2")
	return c.Status(fiber.ErrConflict.Code).JSON(result)
}

func (base *AttachedHandler) Create(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return err
	}

	err = base.service.Create(&object)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(true)

}

func (base *AttachedHandler) Update(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return err
	}

	err = base.service.Update(&object)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(true)
}

func (base *AttachedHandler) Delete(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return err
	}

	err = base.service.Delete(&object)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	return c.Status(fiber.StatusOK).JSON(true)
}

func (base *AttachedHandler) FindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("Id"))
	result, err := base.service.FindById(id)

	if err != nil {
		return useError.ErrorInitServer
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
