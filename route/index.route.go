package route

import (
	"fiber-go/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// cek handler ini
	r.Get("/", handler.UserHandlerRead)

	// route
	r.Get("/users", handler.UserHandlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("/users/:id", handler.UserHandlerGetByID)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Delete("/users/:id", handler.UserHandlerDelete)

}
