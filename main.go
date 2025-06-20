package main

import (
	"todo-api/config"
	"todo-api/models"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Auto-migrate the schema (create table id not exists)
	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8787") // sever starts on http://localhost:8080

}
