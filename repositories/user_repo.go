package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, user.Username, user.Password)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
