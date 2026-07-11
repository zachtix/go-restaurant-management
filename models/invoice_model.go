package model

import "time"

type Invoice struct {
	BaseModel
	Order_id         uint      `json:"order_id"`
	Payment_method   string    `json:"payment_method" validate:"oneof=CARD CASH"`
	Payment_status   string    `json:"payment_status" validate:"oneof=PENDING PAID"`
	Payment_due_date time.Time `json:"payment_due_date"`
}
