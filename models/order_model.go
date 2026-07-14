package model

import "time"

type Order struct {
	BaseModel
	Order_date time.Time `json:"order_date"`
	Table_id   uint      `json:"table_id" validate:"required"`
}
