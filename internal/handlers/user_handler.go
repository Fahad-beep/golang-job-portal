package handlers

import (
	"database/sql"
	"fmt"
	"job_portal/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}
		var userUpdate struct {
			Username string `json:"username"`
			Email string `json:"email"`
		} 
		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")
		fmt.Print(isAdmin)

		if !isAdmin && userID != id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized to update this profile"})
			return
		}

		updateUser, err := services.UpdateUserProfile(db, id, userUpdate.Username, userUpdate.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updateUser)
	}
}
func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}
		user, err := services.GetUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
