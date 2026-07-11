package model

type Table struct {
	BaseModel
	Number_of_guests uint `json:"number_of_guests" validate:"required"`
	Table_Number     uint `json:"table_number" validate:"required"`
}
