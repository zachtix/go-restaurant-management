package controller

import (
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type Controller struct {
	DB *gorm.DB
}

var validate = validator.New()
