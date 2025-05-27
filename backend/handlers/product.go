package handlers

import (
	"clothing-store/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

var products = []models.Product{
	{
		ID:          "1",
		Name:        "Handwoven Cotton Kurta",
		Description: "Traditional Indian kurta made with love and local cotton.",
		Category:    "Men",
		Price:       49.99,
		ImageURL:    "/assets/kurta1.jpg",
		InStock:     true,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "2",
		Name:        "Block Print Saree",
		Description: "Elegant saree with Jaipur-style hand block print.",
		Category:    "Women",
		Price:       89.99,
		ImageURL:    "/assets/saree1.jpg",
		InStock:     true,
		CreatedAt:   time.Now(),
	},
}

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, p := range products {
		if p.ID == id {
			return c.JSON(p)
		}
	}
	return c.Status(404).SendString("Product not found")
}
