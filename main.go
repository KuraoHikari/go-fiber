package main

import (
	"go-fiber/database"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	//     return c.SendString("Hello, World!")
	// })

	routes.RouteInit(app)

	app.Listen(":3000")
}
