package main

import (
	"gorm-crud-app/database"
	"gorm-crud-app/models"
	"gorm-crud-app/handlers"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the User model
	db.AutoMigrate(&models.User{})

	// Set up the router
	router := gin.Default()

	// Define routes
	router.POST("/users", func(c *gin.Context) { handlers.CreateUser(c, db) })
	router.GET("/users", func(c *gin.Context) { handlers.GetUsers(c, db) })
	router.PUT("/users/:id", func(c *gin.Context) { handlers.UpdateUser(c, db) })
	router.DELETE("/users/:id", func(c *gin.Context) { handlers.DeleteUser(c, db) })

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

