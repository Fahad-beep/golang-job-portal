package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserByID(db, id)
}


func UpdateUserProfile(db *sql.DB, id int, userName string, emailID string) (*models.User, error)  {
user := &models.User{ID: id, Username: userName, Email: emailID}
return repository.UpdateUserProfile(db, user)

} 