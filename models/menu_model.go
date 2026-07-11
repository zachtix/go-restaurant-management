package model

import "time"

type Menu struct {
	BaseModel
	Name       string    `json:"name" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
}
