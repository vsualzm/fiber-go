package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fiber-go/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "berhasil mendapatkan api",
		"code":    200,
	})
}

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(fiber.Map{
		"data":    users,
		"message": "success",
	})

}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// validasi
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Debug().Create(&newUser).Error

	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "failed to post data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
		"code":    200,
	})
}
