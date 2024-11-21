package models

import "database/sql"

//import "gorm.io/gorm"

type User struct {
	//gorm.Model
	ID int `json:"id"`
	//Phone    string `json:"phone"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Role       string         `json:"role"`
	City       sql.NullString `json:"city" binding:"omitempty"`
	IsApproved bool           `json:"is_approved"`
}

type Product struct {
	//gorm.Model
	FarmerID    uint     `json:"farmer_id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}
