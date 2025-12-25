package route

import (
	"golang-clean-architecture/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                      *fiber.App
	UserController           *http.UserController
	ContactController        *http.ContactController
	AddressController        *http.AddressController
	AuthCustomerMiddleware   fiber.Handler
	AuthAdminMiddleware      fiber.Handler
	AuthSuperAdminMiddleware fiber.Handler
	AuthDriverMiddleware     fiber.Handler
	AuthMiddleWare           fiber.Handler
	RequestLoggerMiddleware  fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
	c.SetupAuthAdminRoute()
	c.SetupAuthSuperAdminRoute()
	c.SetupAuthDriverRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users/_Login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthAdminRoute() {
	c.App.Use(c.AuthAdminMiddleware)
	// GET /api/driver
	c.App.Get("/api/admin", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Admin API is accessible",
		})
	})

}
func (c *RouteConfig) SetupAuthDriverRoute() {
	c.App.Use(c.AuthDriverMiddleware)

	// GET /api/driver
	c.App.Get("/api/driver", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Driver API is accessible",
		})
	})

}

func (c *RouteConfig) SetupAuthSuperAdminRoute() {
	c.App.Use(c.AuthSuperAdminMiddleware)

	c.App.Get("/api/superadmin", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "superadmin API is accessible",
		})
	})

}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleWare)
	c.App.Use(c.RequestLoggerMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("api/users/_current", c.UserController.Current)

	c.App.Get("/api/contacts", c.ContactController.List)
	c.App.Post("/api/contacts", c.ContactController.Create)
	c.App.Put("/api/contacts/:contactId", c.ContactController.Update)
	c.App.Get("/api/contacts/:contactId", c.ContactController.Get)
	c.App.Delete("/api/contacts/:contactId", c.ContactController.Delete)

	c.App.Get("/api/contacts/:contactId/addresses", c.AddressController.List)
	c.App.Post("/api/contacts/:contactId/addresses", c.AddressController.Create)
	c.App.Put("api/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	c.App.Get("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	c.App.Delete("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)
}
