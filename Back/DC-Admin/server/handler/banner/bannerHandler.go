package banner

import (
	"strconv"

	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/c-go-m/DC-Admin/service/banner"
	"github.com/gofiber/fiber/v2"
)

type BannerHandler struct {
	service banner.IBannerService
}

func New(service banner.IBannerService) *BannerHandler {
	handler := BannerHandler{
		service: service,
	}
	return &handler
}

func (base *BannerHandler) FindAll(c *fiber.Ctx) error {
	result, err := base.service.FindAll()

	if err != nil {
		return useError.ErrorInitServer
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (base *BannerHandler) Create(c *fiber.Ctx) error {
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

func (base *BannerHandler) Update(c *fiber.Ctx) error {
	object, err := base.parceObject(c)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}

	return c.Status(fiber.StatusOK).JSON(base.service.Update(&object))
}

func (base *BannerHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("Id"))
	err := base.service.Delete(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(useError.ControlError(err))
	}
	return c.Status(fiber.StatusOK).JSON(true)
}

func (base *BannerHandler) FindById(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("Id"))
	result, err := base.service.FindById(id)

	if err != nil {
		return useError.ErrorInitServer
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
