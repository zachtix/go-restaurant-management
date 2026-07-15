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

func (h *Controller) GetTables(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	query := h.DB.Model(&model.Table{})

	total, totalPage, err := helper.CountTotal(query, p.Limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var tables []model.Table
	if err := h.DB.Find(&tables).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       tables,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": totalPage,
	})
}
func (h *Controller) GetTable(c fiber.Ctx) error {
	table_id, err := strconv.Atoi(c.Params("table_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var table model.Table
	if err := h.DB.First(&table, "id = ?", table_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "table id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": table})
}
func (h *Controller) CreateTable(c fiber.Ctx) error {
	var table model.Table
	if err := c.Bind().Body(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Create(&table)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": table})
}
func (h *Controller) UpdateTable(c fiber.Ctx) error {
	table_id, err := strconv.Atoi(c.Params("table_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var table model.Table
	if err := c.Bind().Body(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selectedTable model.Table
	if err := h.DB.First(&selectedTable, "id = ?", table_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "table id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	selectedTable.Table_Number = table.Table_Number
	selectedTable.Number_of_guests = table.Number_of_guests

	result := h.DB.Save(&selectedTable)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": selectedTable})
}
func (h *Controller) DeleteTable(c fiber.Ctx) error {
	table_id, err := strconv.Atoi(c.Params("table_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.Table{}, "id = ?", table_id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "table id not found"})
	}
	return c.JSON(fiber.Map{"message": "ok"})
}
