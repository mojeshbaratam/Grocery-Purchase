package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateProduct(db *sql.DB, product *models.Product) error {
	query := "INSERT INTO products (name, price, store_id) VALUES (?, ?, ?)"
	_, err := db.Exec(query, product.Name, product.Price, product.StoreID)
	return err
}

func GetProducts(db *sql.DB) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT p.id, p.name, p.price, p.store_id, s.name as store_name
		FROM products p
		JOIN stores s ON p.store_id = s.id`
	rows, err := db.Query(query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var storeName string
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.StoreID, &storeName); err != nil {
			return products, err
		}

		product.StoreName = storeName
		products = append(products, product)
	}
	return products, nil
}


func GetProductsByName(db *sql.DB, name string) ([]models.Product, error) {
	query := "SELECT p.id, p.name, p.price, s.name as store_name FROM products p JOIN stores s ON p.store_id = s.id WHERE p.name LIKE ?"
	rows, err := db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.StoreName)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(db *sql.DB, name string, newName string, newPrice float64, newStoreId int) (bool, error) {
	query := "UPDATE products SET name = ?, price = ?, store_id = ? WHERE name = ?"
	result, err := db.Exec(query, newName, newPrice, newStoreId, name)
	if err != nil {
		return false, err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	// Return true if the store was updated, false otherwise
	return rowsAffected > 0, nil
}

func RemoveProduct(db *sql.DB, productName string) (bool, error) {
	query := "DELETE FROM products WHERE name = ?"
	result, err := db.Exec(query, productName)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}