package model

type User struct {
	BaseModel
	First_name    string `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     string `json:"last_name" validate:"required,min=2,max=100"`
	Password      string `json:"password" validate:"required,min=6"`
	Email         string `json:"email" gorm:"unique" validate:"email,required"`
	Avatar        string `json:"avatar"`
	Phone         string `json:"phone" validate:"required"`
	Refresh_Token string `json:"refresh_token"`
}
