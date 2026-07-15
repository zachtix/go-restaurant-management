package controller

import (
	"errors"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetOrders(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	tableID := fiber.Query[int](c, "table_id")
	open := fiber.Query[bool](c, "open")

	query := h.DB.Model(&model.Order{})
	if tableID > 0 {
		query = query.Where("table_id = ?", tableID)
	}
	if open {
		query = query.Where("id NOT IN (?)",
			h.DB.Model(&model.Invoice{}).Select("order_id"),
		)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var orders []model.Order
	if err := query.Limit(p.Limit).Offset(p.Offset).Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       orders,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": (total + int64(p.Limit) - 1) / int64(p.Limit),
	})
}
func (h *Controller) GetOrder(c fiber.Ctx) error {
	order_id, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order model.Order
	if err := h.DB.First(&order, "id = ?", order_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "order not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": order})
}
func (h *Controller) CreateOrder(c fiber.Ctx) error {
	var order model.Order
	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	order.Order_date = time.Now()
	if err := validate.Struct(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Create(&order)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": order})
}
func (h *Controller) UpdateOrder(c fiber.Ctx) error {
	order_id, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selectedOrder model.Order
	if err := h.DB.First(&selectedOrder, "id = ?", order_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "order id not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order model.Order
	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	selectedOrder.Table_id = order.Table_id

	result := h.DB.Save(&selectedOrder)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": selectedOrder})
}
func (h *Controller) DeleteOrder(c fiber.Ctx) error {
	order_id, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.Order{}, order_id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "order id not found"})
	}
	return c.JSON(fiber.Map{"message": "ok"})
}
