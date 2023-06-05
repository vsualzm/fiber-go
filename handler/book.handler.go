package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fiber-go/model/request"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(c *fiber.Ctx) error {

	// buat request baru
	book := new(request.BookCreateRequest)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	// validasi
	validate := validator.New()
	errValidate := validate.Struct(book)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// // handle file
	// file, errFile := c.FormFile("cover")
	// if errFile != nil {
	// 	log.Println("error file = ", errFile)
	// }
	// // cek Skip error upload_file
	// var filename string
	// if file != nil {
	// 	filename = file.Filename
	// 	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", file.Filename))
	// 	if errSaveFile != nil {
	// 		log.Println("fail to store file into public/covers directory")
	// 	}
	// } else {
	// 	log.Println("Nothing file to uploading")
	// }

	// validation required image
	var filenameString string

	filename := c.Locals("filename")
	log.Println("filename = ", filename)

	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "Image cover is reqired.",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)

	}

	newBook := entity.Book{
		Author:   book.Author,
		Title:    book.Title,
		Cover:    filenameString,
		UpdateAt: time.Now(),
	}

	// add to database
	errCreateDatabase := database.DB.Debug().Create(&newBook).Error

	// jika gagal
	if errCreateDatabase != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to post data",
		})
	}

	// jika status success

	return c.Status(200).JSON(fiber.Map{
		"message": "succes",
		"data":    newBook,
		"code":    "200",
	})

}

func GetAllBook(c *fiber.Ctx) error {
	// generate cek info login user admin
	userInfo := c.Locals("userInfo")
	log.Println("user info data:", userInfo)

	var books []entity.Book

	result := database.DB.Debug().Find(&books)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(fiber.Map{
		"data":    books,
		"message": "success",
	})

}
