package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fiber-go/model/request"
	"fiber-go/model/response"
	"fiber-go/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GETALL
func UserHandlerGetAll(c *fiber.Ctx) error {

	// generate cek info login user admin
	userInfo := c.Locals("userInfo")
	log.Println("user info data:", userInfo)

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

// POST
func UserHandlerCreate(c *fiber.Ctx) error {

	// membuat request baru
	user := new(request.UserCreateRequest)

	// memasukan ke bodyparser
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

	// mendapatkan yg user masukan
	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		UpdateAt: time.Now(),
	}

	// hashing password

	hashedPassword, err := utils.HashingPassword(user.Password)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	// memasukan data ke dalam database
	errCreateUser := database.DB.Debug().Create(&newUser).Error

	// jika gagal berikan post ini
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to post data",
		})
	}

	// jika status succes kembalikan status ini
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
		"code":    200,
	})

}

func UserHandlerGetByID(c *fiber.Ctx) error {

	// ambil id
	userID := c.Params("ID")

	// inisiasi user ke dalam struct
	var user entity.User

	// cari di database id yg sama
	err := database.DB.First(&user, "ID = ?", userID).Error

	// response error ketika  data tidak di temukan
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	// mapping response suapaya response di kembalikan tidak ada nilai delete
	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}

	// response success
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserHandlerUpdate(c *fiber.Ctx) error {

	userRequest := new(request.UserUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	// inisiasi user ke dalam struct
	var user entity.User
	// ambil id
	userID := c.Params("ID")
	// cek data di database
	err := database.DB.First(&user, "ID = ?", userID).Error
	// response error ketika  data tidak di temukan
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	// update user data
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	user.UpdateAt = time.Now()

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	// response success
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(c *fiber.Ctx) error {

	// ini untuk update email
	userRequest := new(request.UserEmailRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	// inisiasi user ke dalam struct
	var user entity.User
	var isEmailUserExist entity.User

	// ambil id
	userID := c.Params("ID")

	// cek Availbel user
	err := database.DB.First(&user, "ID = ?", userID).Error
	// response error ketika  data tidak di temukan
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	// cek Availbel Email
	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error

	if errCheckEmail == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Email alredy to use",
		})
	}

	// update user data
	user.Email = userRequest.Email
	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return c.Status(402).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	// response success
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(c *fiber.Ctx) error {
	// dapatkan ID
	userID := c.Params("ID")

	// beri variabel user
	var user entity.User

	// cari ID yg sama
	err := database.DB.Debug().First(&user, "ID = ? ", userID).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not foudn",
		})
	}

	// ambil ke database dan delete id yg di ingin kan
	errDelete := database.DB.Debug().Delete(&user).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user was deleted",
	})

}

// cek rest api
func UserHandlerRead(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "berhasil mendapatkan api",
		"code":    200,
	})
}
