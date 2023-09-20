package server

import (
	"github.com/c-go-m/DC-Admin/common/utility/config"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
	"github.com/c-go-m/DC-Admin/server/handler"
	"github.com/c-go-m/DC-Admin/server/handler/attached"
	"github.com/c-go-m/DC-Admin/server/handler/banner"
	"github.com/c-go-m/DC-Admin/server/handler/category"
	"github.com/c-go-m/DC-Admin/server/handler/information"
	"github.com/c-go-m/DC-Admin/server/handler/menu"
	"github.com/c-go-m/DC-Admin/server/handler/presentation"
	"github.com/c-go-m/DC-Admin/server/handler/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	app *fiber.App
}

func New() *Server {
	return &Server{}
}

func (base *Server) InitServer() (err error) {
	initDependency()

	base.app = fiber.New(getConfig())

	base.app.Use(validateAccess)
	base.app.Use(controlError)
	base.app.Use(cors.New(getConfigCors()))

	base.AddHandlers()

	if err = base.app.Listen(config.Port); err != nil {
		return useError.ErrorInitServer
	}

	return
}

func getConfig() fiber.Config {

	fiber_config := fiber.Config{
		ServerHeader:  config.ServerHeader,
		StrictRouting: config.StrictRouting,
	}

	return fiber_config
}

func (base *Server) StopServer() (err error) {
	if err = base.app.Shutdown(); err != nil {
		return useError.ErrorInitServer
	}
	return
}

func (base *Server) AddHandlers() {
	base.HealthCheck()
	base.AddListHandlers(banner.New(bannerService))
	base.AddListHandlers(attached.New(attachedService))
	base.AddListHandlers(category.New(categoryService))
	base.AddListHandlers(menu.New(menuService))
	base.AddListHandlers(product.New(productService))
	base.AddListHandlers(presentation.New(presentationService))
	base.AddListHandlers(information.New(informationService))
}

func (base *Server) HealthCheck() {
	base.app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("HealthCheck")
	})
}

func (base *Server) AddListHandlers(handler handler.IBaseHandler) {

	for k, v := range handler.AddHandlersGet() {
		base.app.Get(k, v)
	}

	for k, v := range handler.AddHandlersPost() {
		base.app.Post(k, v)
	}

	for k, v := range handler.AddHandlersPut() {
		base.app.Put(k, v)
	}

	for k, v := range handler.AddHandlersDelete() {
		base.app.Delete(k, v)
	}
}
