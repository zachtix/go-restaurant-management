package model

type OrderItem struct {
	BaseModel
	Quantity   string  `json:"quantity" validate:"oneof=S M L"`
	Unit_price float64 `json:"unit_price"`
	Food_id    uint    `json:"food_id" validate:"required"`
	Order_id   uint    `json:"order_id" validate:"required"`
}
