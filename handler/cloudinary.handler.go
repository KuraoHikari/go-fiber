package handler

import (
	"context"
	"go-fiber/model/entity"
	"log"

	"go-fiber/database"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

// cloudinaryURL := os.Getenv("DB_NAME")

// cldUrl := fmt.Sprintf("%s", cloudinaryURL)

func Clodinary(c *fiber.Ctx) error {
	cloudinaryURL := "cloudinary://187461791293518:k5dLmzTgI8j_iOxvKOgU_tmznjc@dw0nynnmv"

	var clodinaryData entity.Cloudinary

	_ = c.BodyParser(&clodinaryData)

	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()
	// log.Println(file.Filename)

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(cloudinaryURL)

	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(resp.SecureURL)

	clodinaryData.Image = resp.SecureURL

	database.DB.Create(&clodinaryData)

	return c.JSON(fiber.Map{
		"data": clodinaryData,
	})
}
