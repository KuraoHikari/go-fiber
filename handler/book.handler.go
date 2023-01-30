package handler

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(book)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed create book",
			"error":   errValidate.Error(),
		})
	}

	filename := ctx.Locals("filename")
	var filenameString string

	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
