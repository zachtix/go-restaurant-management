package controller

import (
	"errors"
	model "restaurant-management/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetMenus(c fiber.Ctx) error                { return c.JSON(fiber.Map{"message": "ok"}) }
func (h *Controller) GetMenu(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateMenu(c fiber.Ctx) error              { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateMenu(c fiber.Ctx) error              { return c.JSON(fiber.Map{"message": "ok"}) }

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
