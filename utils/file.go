package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	//filenya
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("error file = ", errFile)
		// return ctx.Status(500).JSON(fiber.Map{
		// 	"message": "failed to store data",
		// })
	}
	var filename *string
	if file != nil {
		filename = &file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))
		if errSaveFile != nil {
			log.Println("filed to stored file")
			// return ctx.Status(500).JSON(fiber.Map{
			// 	"message": "failed to store data",
			// })
		}
	} else {
		log.Println("Nothing file to uploading")
	}
	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}
