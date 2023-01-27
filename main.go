package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

// var (
// 	db *gorm.DB = database.SetupDatabaseConnection()
// )

func main() {
	defer database.CloseDatabaseConnection(database.DB)
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":3000")
}
