package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	InStock     bool      `json:"in_stock"`
	CreatedAt   time.Time `json:"created_at"`
}
