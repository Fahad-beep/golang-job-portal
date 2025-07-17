package handlers

import (
	"database/sql"
	"fmt"
	"job_portal/internal/models"
	"job_portal/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := services.LoginHandler(db, user.Username, user.Password)
		// log.Fatal(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password or username"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(user)
		err := services.RegisterUser(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Succesfully created user"})

	}
}
