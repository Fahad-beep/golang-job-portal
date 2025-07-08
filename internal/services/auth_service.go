package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
	"job_portal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	hashedPws, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPws)
	return repository.CreateUser(db, user)
}

func LoginHandler(db *sql.DB, username, password string) (string, error) {
	user, err := repository.GetUserByUserName(db, username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// log.Fatal(err)
		return "", err
	}
	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}
