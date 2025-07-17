package routes

import (
	"database/sql"
	"job_portal/internal/auth"
	"job_portal/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {

	//global routes
	r.POST("/login", handlers.LoginHandler(db))
	r.POST("/register", handlers.RegisterHandler(db))

	//user routes
	// authenticated := r.Group("/")
	// authenticated.Use(auth.AuthMiddleware()) 
	r.GET("/user/:id", auth.AuthMiddleware(), handlers.GetUserByIdHandler(db))
	r.PUT("/user/:id", auth.AuthMiddleware(), handlers.UpdateUserProfileHandler(db))
}
