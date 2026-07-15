package controller

import (
	"errors"
	"restaurant-management/helper"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetInvoices(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	query := h.DB.Model(&model.Invoice{})

	total, totalPage, err := helper.CountTotal(query, p.Limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var invoices []model.Invoice
	if err := query.Limit(p.Limit).Offset(p.Offset).Find(&invoices).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       invoices,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": totalPage,
	})
}

func (h *Controller) GetInvoice(c fiber.Ctx) error {
	inovice_id, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var invoice model.Invoice
	if err := h.DB.First(&invoice, "id = ?", inovice_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "invoice id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var order model.Order
	if err := h.DB.First(&order, "id = ?", invoice.Order_id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "order id not found"})
	}

	var order_items []model.OrderItem
	if err := h.DB.Find(&order_items, "order_id = ?", invoice.Order_id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "order item by order id not found"})
	}
	var total float64
	for _, e := range order_items {
		total += e.Unit_price
	}

	return c.JSON(fiber.Map{
		"message": "ok",
		"data": fiber.Map{
			"invoice_id":       invoice.ID,
			"order_id":         invoice.Order_id,
			"table_id":         order.Table_id, // order ที่ query มาตอน validate
			"payment_method":   invoice.Payment_method,
			"payment_status":   invoice.Payment_status,
			"payment_due_date": invoice.Payment_due_date,
			"total_amount":     total,
			"order_items":      order_items,
		},
	})
}

func (h *Controller) CreateInvoice(c fiber.Ctx) error {

	var invoice model.Invoice
	if err := c.Bind().Body(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var existing model.Invoice
	if err := h.DB.Where("order_id = ?", invoice.Order_id).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "invoice already exists"})
	}
	invoice.Payment_status = "PENDING"
	if err := validate.Struct(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order model.Order
	if err := h.DB.First(&order, "id = ?", invoice.Order_id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "order id not found"})
	}

	var order_items []model.OrderItem
	if err := h.DB.Find(&order_items, "order_id = ?", invoice.Order_id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "order item by order id not found"})
	}
	var total float64
	for _, e := range order_items {
		total += e.Unit_price
	}

	invoice.Payment_due_date = time.Now().Add(time.Hour * 1)

	result := h.DB.Create(&invoice)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
		"data": fiber.Map{
			"invoice_id":       invoice.ID,
			"order_id":         invoice.Order_id,
			"table_id":         order.Table_id, // order ที่ query มาตอน validate
			"payment_method":   invoice.Payment_method,
			"payment_status":   invoice.Payment_status,
			"payment_due_date": invoice.Payment_due_date,
			"total_amount":     total,
			"order_items":      order_items,
		},
	})
}

func (h *Controller) UpdateInvoice(c fiber.Ctx) error {
	inovice_id, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var invoice model.Invoice
	if err := c.Bind().Body(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selected_invoice model.Invoice
	if err := h.DB.First(&selected_invoice, "id = ?", inovice_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "invoice id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if selected_invoice.Payment_status == "PAID" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "invoice already paid"})
	}

	selected_invoice.Payment_method = invoice.Payment_method
	selected_invoice.Payment_status = invoice.Payment_status

	result := h.DB.Save(&selected_invoice)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok"})
}

func (h *Controller) DeleteInvoice(c fiber.Ctx) error {
	inovice_id, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.Invoice{}, inovice_id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "invoice id not found"})
	}
	return c.JSON(fiber.Map{"message": "ok"})
}
