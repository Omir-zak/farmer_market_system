package models

//import "gorm.io/gorm"

type User struct {
	//gorm.Model
	ID int `json:"id"`
	//Phone    string `json:"phone"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Role       string  `json:"role"`
	City       *string `json:"city,omitempty"`
	IsApproved bool    `json:"is_approved"`
}
