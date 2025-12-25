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
	admin := c.App.Group("/api/admin", c.AuthAdminMiddleware)
	admin.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Admin API is accessible",
		})
	})

}
func (c *RouteConfig) SetupAuthDriverRoute() {
	driver := c.App.Group("/api/driver", c.AuthDriverMiddleware)
	driver.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Driver API is accessible",
		})
	})

}

func (c *RouteConfig) SetupAuthSuperAdminRoute() {
	superAdmin := c.App.Group("/api/superadmin", c.AuthSuperAdminMiddleware)
	superAdmin.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Superadmin API is accessible",
		})
	})

}

func (c *RouteConfig) SetupAuthRoute() {
	admin := c.App.Group("/api/admin", c.AuthAdminMiddleware)

	admin.Delete("/api/users", c.UserController.Logout)
	admin.Patch("/api/users/_current", c.UserController.Update)
	admin.Get("api/users/_current", c.UserController.Current)

	admin.Get("/api/contacts", c.ContactController.List)
	admin.Post("/api/contacts", c.ContactController.Create)
	admin.Put("/api/contacts/:contactId", c.ContactController.Update)
	admin.Get("/api/contacts/:contactId", c.ContactController.Get)
	admin.Delete("/api/contacts/:contactId", c.ContactController.Delete)

	admin.Get("/api/contacts/:contactId/addresses", c.AddressController.List)
	admin.Post("/api/contacts/:contactId/addresses", c.AddressController.Create)
	admin.Put("api/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	admin.Get("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	admin.Delete("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)
}
