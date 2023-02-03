package main

import (
	"fiber-go/database"
	"fiber-go/database/migration"
	"fiber-go/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	// RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":3000")
}
