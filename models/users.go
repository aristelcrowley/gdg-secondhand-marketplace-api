package models

type User struct {
	ID       int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
