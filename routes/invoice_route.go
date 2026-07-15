package route

import (
	controller "restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func InvoiceRoute(app *fiber.App, h *controller.Controller) {
	invoice := app.Group("/invoices")
	invoice.Use(middleware.Authenticate)
	invoice.Get("", middleware.Paginate, h.GetInvoices)
	invoice.Get("/:invoice_id", h.GetInvoice)
	invoice.Post("", h.CreateInvoice)
	invoice.Patch("/:invoice_id", h.UpdateInvoice)
	invoice.Delete("/:invoice_id", h.DeleteInvoice)
}
