package controller

import (
	"errors"
	"restaurant-management/helper"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetOrderItems(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	query := h.DB.Model(&model.OrderItem{})

	total, totalPage, err := helper.CountTotal(query, p.Limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order_items []model.OrderItem
	if err := h.DB.Limit(p.Limit).Offset(p.Offset).Find(&order_items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       order_items,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": totalPage,
	})
}

func (h *Controller) GetOrderItem(c fiber.Ctx) error {
	order_item_id, err := strconv.Atoi(c.Params("order_item_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order_item model.OrderItem
	if err := h.DB.First(&order_item, "id = ?", order_item_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": order_item})
}

func (h *Controller) GetOrderItemByOrder(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	query := h.DB.Model(&model.OrderItem{})

	total, totalPage, err := helper.CountTotal(query, p.Limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	order_id, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order_items []model.OrderItem
	if err := query.Limit(p.Limit).Offset(p.Offset).Find(&order_items, "order_id = ?", order_id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       order_items,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": totalPage,
	})
}

func (h *Controller) CreateOrderItem(c fiber.Ctx) error {
	var order_item model.OrderItem
	if err := c.Bind().Body(&order_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var order model.Order
	if err := h.DB.First(&order, "id = ?", order_item.Order_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "order id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var food model.Food
	if err := h.DB.First(&food, "id = ?", order_item.Food_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "food id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate.Struct(&order_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	order_item.Unit_price = food.Price
	result := h.DB.Create(&order_item)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": order_item})
}

func (h *Controller) UpdateOrderItem(c fiber.Ctx) error {
	order_item_id, err := strconv.Atoi(c.Params("order_item_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var order_item model.OrderItem
	if err := c.Bind().Body(&order_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&order_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selected_order_item model.OrderItem
	if err := h.DB.First(&selected_order_item, "id = ?", order_item_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	selected_order_item.Quantity = order_item.Quantity

	result := h.DB.Save(&selected_order_item)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": selected_order_item})
}

func (h *Controller) DeleteOrderItem(c fiber.Ctx) error {
	order_item_id, err := strconv.Atoi(c.Params("order_item_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.OrderItem{}, order_item_id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "order item id not found"})
	}
	return c.JSON(fiber.Map{"message": "ok"})
}
