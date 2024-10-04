package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateAdmin(db *sql.DB, admin *models.Admin) error {
	query := "INSERT INTO admins (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, admin.Username, admin.Password)
	return err
}

func GetAdminByUsername(db *sql.DB, username string) (models.Admin, error) {
	var admin models.Admin
	query := "SELECT * FROM admins WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&admin.ID, &admin.Username, &admin.Password)
	return admin, err
}
