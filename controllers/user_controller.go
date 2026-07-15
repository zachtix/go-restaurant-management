package controller

import (
	"errors"
	"os"
	"restaurant-management/helper"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Controller) GetUsers(c fiber.Ctx) error {
	p := middleware.GetPagination(c)

	query := h.DB.Model(&model.User{})

	total, totalPage, err := helper.CountTotal(query, p.Limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var users []model.User
	if err := h.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       users,
		"page":       p.Page,
		"limit":      p.Limit,
		"total":      total,
		"total_page": totalPage,
	})
}
func (h *Controller) GetUser(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var user model.User
	if err := h.DB.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user id not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "ok", "data": user})
}
func (h *Controller) CreateUser(c fiber.Ctx) error {
	var user model.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	user.Password = hashPassword

	result := h.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "ok"})
}
func (h *Controller) LoginUser(c fiber.Ctx) error {
	var user model.User
	var selectUser model.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	result := h.DB.Where("email = ?", user.Email).First(&selectUser)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error})
	}

	err := helper.ComparePassword(selectUser.Password, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Wrong Email or Password"})
	}

	access, err := helper.GenerateJWT(selectUser, os.Getenv("SECRET_ACCESS"), 1)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Generate JWT Error"})
	}
	refresh, err := helper.GenerateJWT(selectUser, os.Getenv("SECRET_REFRESH"), 240)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Generate JWT Error"})
	}

	return c.JSON(fiber.Map{"message": "ok", "access_token": access, "refresh_token": refresh})
}
