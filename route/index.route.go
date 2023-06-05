package route

import (
	"fiber-go/config"
	"fiber-go/handler"
	"fiber-go/middleware"
	"fiber-go/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// cek handler ini
	r.Get("/", handler.UserHandlerRead)

	// static asset
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	// auth
	r.Post("/login", handler.LoginHandler)

	// route - CRUD
	r.Get("/users", middleware.Auth, handler.UserHandlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("/users/:id", handler.UserHandlerGetByID)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Put("/users/:id/email-update", handler.UserHandlerUpdateEmail)
	r.Delete("/users/:id", handler.UserHandlerDelete)

	// Route - CRUD book
	// middleware.Auth
	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
	r.Get("/book", handler.GetAllBook)
	// r.Get("/book/:id", handler.UserHandlerGetByID)
	// r.Put("/book/:id", handler.UserHandlerUpdate)
	// r.Delete("/book/:id", handler.UserHandlerDelete)

}
