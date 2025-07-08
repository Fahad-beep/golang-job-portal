package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func CraeteUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`INSERT INTO users (username, password, email) VALUES (?, ?, ?)`, user.Username, user.Password, user.Email)
	return err
}
