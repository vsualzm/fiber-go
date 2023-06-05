package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	// handle file
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("error file = ", errFile)
	}

	// cek Skip error upload_file
	var filename *string

	if file != nil {
		filename = &file.Filename
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))
		if errSaveFile != nil {
			log.Println("fail to store file into public/covers directory")
		}
	} else {
		log.Println("Nothing file to uploading")
	}

	if filename != nil {
		c.Locals("filename", *filename)
	} else {
		c.Locals("filename", nil)
	}
	return c.Next()

}
