package models

type Movie struct {
	ID        string `gorm:"primary_key;not null;unique" json:"id"`
	MovieName string `json:"moviename"`
}
type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}
