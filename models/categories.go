package models

type Category struct {
	ID   int    `json:"category_id" gorm:"primaryKey;column:category_id"`
	Name string `json:"category_name" gorm:"column:category_name"`
}
