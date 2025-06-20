package main

import (
	"todo-api/models"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
	"honnef.co/go/tools/config"
)

func main() {
	config.ConnectDatabase()

	// Auto-migrate the schema (create table id not exists)
	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080") // sever starts on http://localhost:8080

}
