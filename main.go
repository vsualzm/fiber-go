package main

import (
	"fiber-go/database"
	"fiber-go/database/migration"
	"fiber-go/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// INISIASI DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	// INISIASI ROUTE
	route.RouteInit(app)

	// port API
	log.Fatal(app.Listen(":1234"))
}
