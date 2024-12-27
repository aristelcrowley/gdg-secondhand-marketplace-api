package models

type User struct {
	ID       int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	AuthToken string `json:"auth_token" gorm:"column:auth_token"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
