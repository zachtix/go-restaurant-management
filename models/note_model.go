package model

type Note struct {
	BaseModel
	Text  string `json:"text"`
	Title string `json:"title"`
}
