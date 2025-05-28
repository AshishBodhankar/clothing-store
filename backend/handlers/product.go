package handlers

import (
	"clothing-store/backend/models"

	"github.com/gofiber/fiber/v2"

	"strconv"
)

// var products = []models.Product{
// 	{
// 		ID:          1,
// 		Name:        "Handwoven Cotton Kurta",
// 		Description: "Traditional Indian kurta made with love and local cotton.",
// 		Category:    "Men",
// 		Price:       49.99,
// 		ImageURL:    "/assets/kurta1.jpg",
// 		InStock:     true,
// 		CreatedAt:   time.Now(),
// 	},
// 	{
// 		ID:          2,
// 		Name:        "Block Print Saree",
// 		Description: "Elegant saree with Jaipur-style hand block print.",
// 		Category:    "Women",
// 		Price:       89.99,
// 		ImageURL:    "/assets/saree1.jpg",
// 		InStock:     true,
// 		CreatedAt:   time.Now(),
// 	},
// }

func GetAllProducts(c *fiber.Ctx) error {
	products, err := models.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	product, err := models.GetProductByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(product)
}
