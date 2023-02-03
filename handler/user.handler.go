package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fiber-go/model/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result != nil {
		log.Println(result.Error)
	}

	// err := database.DB(&users).Error
	// if err != nil {
	// 	log.Println(err)
	// }
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.JSON(users)
	}

}
