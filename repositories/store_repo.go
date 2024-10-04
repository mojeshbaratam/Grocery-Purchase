package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateStore(db *sql.DB, store *models.Store) error {
	query := "INSERT INTO stores (name) VALUES (?)"
	_, err := db.Exec(query, store.Name)
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
		if err := rows.Scan(&store.ID, &store.Name); err != nil {
			return stores, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}
