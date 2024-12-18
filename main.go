package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-gin-postgres/config"
	"go-gin-postgres/controllers"
	"go-gin-postgres/routes"
	"log"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	db := config.ConnectDB()
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Initialize controllers
	todoController := controllers.NewTodoController(db)

	// Setup routes
	routes.SetupTodoRoutes(router, todoController)

	// Run the server
	router.Run(":8080")
}
