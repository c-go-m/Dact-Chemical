package handler

import "github.com/gofiber/fiber/v2"

type IBaseHandler interface {
	AddHandlersGet() map[string]func(*fiber.Ctx) error
	AddHandlersPost() map[string]func(*fiber.Ctx) error
	AddHandlersPut() map[string]func(*fiber.Ctx) error
	AddHandlersDelete() map[string]func(*fiber.Ctx) error
}
