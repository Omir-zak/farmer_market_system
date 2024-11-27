package controllers

import (
	"database/sql"
	"farmer_market/models"
	"github.com/lib/pq"
)

// Получение продуктов фермера
func GetFarmerProducts(db *sql.DB, farmerID int) ([]models.Product, error) {
	query := `SELECT id, name, category, price, quantity, description, images, is_out_of_stock
              FROM products WHERE farmer_id = $1`
	rows, err := db.Query(query, farmerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price,
			&product.Quantity, &product.Description, pq.Array(&product.Images), &product.IsOutOfStock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Добавление нового продукта
func AddProduct(db *sql.DB, product *models.Product) error {
	query := `INSERT INTO products (farmer_id, name, category, price, quantity, description, images, is_out_of_stock)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(query, product.FarmerID, product.Name, product.Category, product.Price,
		product.Quantity, product.Description, pq.Array(product.Images), product.IsOutOfStock)
	return err
}

// Обновление продукта
func UpdateProduct(db *sql.DB, productID int, product *models.Product) error {
	query := `UPDATE products SET name = $1, category = $2, price = $3, quantity = $4, 
              description = $5, images = $6, is_out_of_stock = $7
              WHERE id = $8 AND farmer_id = $9`
	_, err := db.Exec(query, product.Name, product.Category, product.Price, product.Quantity,
		product.Description, pq.Array(product.Images), product.IsOutOfStock, productID, product.FarmerID)
	return err
}

// Удаление продукта
func DeleteProduct(db *sql.DB, productID, farmerID int) error {
	query := `DELETE FROM products WHERE id = $1 AND farmer_id = $2`
	_, err := db.Exec(query, productID, farmerID)
	return err
}
