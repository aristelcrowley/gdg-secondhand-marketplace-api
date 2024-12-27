package models

type Item struct {
	ID          int    `json:"item_id" gorm:"primaryKey;column:item_id"`
	UserID      int    `json:"user_id" gorm:"column:user_id"`
	CategoryID  int    `json:"category_id" gorm:"column:category_id"`
	Name        string `json:"item_name" gorm:"column:item_name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stok" gorm:"column:stok"`
	Description string `json:"deskripsi" gorm:"column:deskripsi"`
}
