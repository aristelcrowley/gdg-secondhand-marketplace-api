package models

type Order struct {
	ID         int `json:"order_id" gorm:"primaryKey;column:order_id"`
	UserID     int `json:"user_id" gorm:"column:user_id"`
	ItemID     int `json:"item_id" gorm:"column:item_id"`
	ItemAmount int `json:"item_amount"`
	PriceTotal int `json:"price_total"`
}
