package controller

import (
	"errors"
	model "restaurant-management/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetOrders(c fiber.Ctx) error {
	var orders []model.Order
	if err := h.DB.Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": orders})
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

	selectedOrder.Order_date = order.Order_date
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
