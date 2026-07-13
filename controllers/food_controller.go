package controller

import (
	"errors"
	model "restaurant-management/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var validate = validator.New()

func (h *Controller) GetFoods(c fiber.Ctx) error {
	var foods []model.Food
	result := h.DB.Find(&foods)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": foods})
}

func (h *Controller) GetFood(c fiber.Ctx) error {
	food_id, err := strconv.Atoi(c.Params("food_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var food model.Food
	if err := h.DB.First(&food, "id = ?", food_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "food not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": food})
}

func (h *Controller) CreateFood(c fiber.Ctx) error {
	var food model.Food
	if err := c.Bind().Body(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Create(&food)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "ok", "data": food})
}

func (h *Controller) UpdateFood(c fiber.Ctx) error {
	food_id, err := strconv.Atoi(c.Params("food_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var food model.Food
	if err := c.Bind().Body(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Struct(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var selectedFood model.Food
	result := h.DB.Where("id = ?", food_id).First(&selectedFood)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": result.Error.Error()})
	}
	selectedFood.Name = food.Name
	selectedFood.Price = food.Price
	selectedFood.Food_image = food.Food_image
	selectedFood.Menu_id = food.Menu_id

	save := h.DB.Save(&selectedFood)
	if save.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": save.Error.Error()})
	}
	return c.JSON(fiber.Map{"message": "ok", "data": selectedFood})
}

func (h *Controller) DeleteFood(c fiber.Ctx) error {
	food_id, err := strconv.Atoi(c.Params("food_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result := h.DB.Delete(&model.Food{}, food_id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "food not found"})
	}

	return c.JSON(fiber.Map{"message": "ok"})
}
