package route

import (
	controller "restaurant-management/controllers"

	"github.com/gofiber/fiber/v3"
)

func InvoiceRoute(app *fiber.App, h *controller.Controller) {
	invoice := app.Group("/invoices")
	invoice.Get("", h.GetInvoices)
	invoice.Get("/:invoice_id", h.GetInvoice)
	invoice.Post("", h.CreateInvoice)
	invoice.Patch("/:invoice_id", h.UpdateInvoice)
	invoice.Delete("/:invoice_id", h.DeleteInvoice)
}
