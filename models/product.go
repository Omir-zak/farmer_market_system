package models

type Product struct {
	ID           int      `json:"id"`
	FarmerID     int      `json:"farmer_id"`
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Price        float64  `json:"price"`
	Quantity     int      `json:"quantity"`
	Description  string   `json:"description"`
	Images       []string `json:"images"`
	IsOutOfStock bool     `json:"is_out_of_stock"`
}
