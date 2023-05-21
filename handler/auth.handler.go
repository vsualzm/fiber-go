package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fiber-go/model/request"
	"fiber-go/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(c *fiber.Ctx) error {
	//1. membuat variabel request
	loginRequest := new(request.LoginRequest)

	//2. memasukan ke bodyparser
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	// 3. validasi
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// 4. inisiasi user ke dalam struct
	var user entity.User

	// 5. cari di database id yg avaible
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error

	// 6. response error ketika  data tidak di temukan
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Error Credential",
		})
	}

	// 7. cek validasi hash
	isValid := utils.CheckHash(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	// 8. generated TOKEN

	// yang di kirim jwt berbentuk tokennya
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		// debug
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//9.  ketika success
	return c.JSON(fiber.Map{
		"token": token,
	})
}
