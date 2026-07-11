package model

type Food struct {
	BaseModel
	Name       string  `json:"name" validate:"required,min=2,max=100"`
	Price      float64 `json:"price" validate:"required,numeric,min=1"`
	Food_image string  `json:"food_image" validate:"required"`
	Menu_id    uint    `json:"menu_id" validate:"required"`
}
