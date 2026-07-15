package controller

import (
	"errors"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetMenus(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	var total int64
	if err := h.DB.Model(&model.Menu{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var menus []model.Menu
	if err := h.DB.Limit(p.Limit).Offset(p.Offset).Find(&menus).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       menus,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": (total + int64(p.Limit) - 1) / int64(p.Limit),
	})
}
func (h *Controller) GetMenu(c fiber.Ctx) error {
	menu_id, err := strconv.Atoi(c.Params("menu_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var menu model.Menu
	if err := h.DB.First(&menu, "id = ?", menu_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": menu})
}
func (h *Controller) CreateMenu(c fiber.Ctx) error {
	var menu model.Menu
	if err := c.Bind().Body(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Create(&menu)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "ok", "data": menu})
}
func (h *Controller) UpdateMenu(c fiber.Ctx) error {
	menu_id, err := strconv.Atoi(c.Params("menu_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var menu model.Menu
	if err := c.Bind().Body(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selectedMenu model.Menu
	if err := h.DB.First(&selectedMenu, "id = ?", menu_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "menu id not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	selectedMenu.Name = menu.Name
	selectedMenu.Category = menu.Category
	selectedMenu.Start_date = menu.Start_date
	selectedMenu.End_date = menu.End_date

	result := h.DB.Save(&selectedMenu)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": menu})
}
func (h *Controller) DeleteMenu(c fiber.Ctx) error {
	menu_id, err := strconv.Atoi(c.Params("menu_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.Menu{}, menu_id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "menu not found"})
	}
	return c.JSON(fiber.Map{"message": "ok"})
}

func (h *Controller) MenuIDExists(menu_id uint) (bool, error) {
	var menu model.Menu
	if err := h.DB.First(&menu, "id = ?", menu_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
