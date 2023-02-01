package main

import (
	"fiber-go/database"
	"fiber-go/database/migration"
	"fiber-go/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial database
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
