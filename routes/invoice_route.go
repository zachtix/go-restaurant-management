package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func InvoiceRoute(app *fiber.App) {
	invoice := app.Group("/invoices")
	invoice.Get("", controller.GetInvoices)
	invoice.Get("/:invoice_id", controller.GetInvoice)
	invoice.Post("", controller.CreateInvoice)
	invoice.Patch("/:invoice_id", controller.UpdateInvoice)
}
