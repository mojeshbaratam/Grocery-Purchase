package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateStore(db *sql.DB, store *models.Store) error {
	query := "INSERT INTO stores (name, location) VALUES (?, ?)"
	_, err := db.Exec(query, store.Name, store.Location)
	return err
}

func GetStores(db *sql.DB) ([]models.Store, error) {
	var stores []models.Store
	query := "SELECT * FROM stores"
	rows, err := db.Query(query)
	if err != nil {
		return stores, err
	}
	defer rows.Close()

	for rows.Next() {
		var store models.Store
		if err := rows.Scan(&store.ID, &store.Name, &store.Location); err != nil {
			return stores, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}

func UpdateStore(db *sql.DB, name string, newName string, newLocation string) (bool, error) {
	query := "UPDATE stores SET name = ?, location = ? WHERE name = ?"
	result, err := db.Exec(query, newName, newLocation, name)
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

func RemoveStore(db *sql.DB, storeName string) (bool, error) {
	query := "DELETE FROM stores WHERE name = ?"
	result, err := db.Exec(query, storeName)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}