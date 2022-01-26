package rest

import (
	file2 "github.com/ciazhar/emobi-service/third_party/file"
	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/gofiber/fiber/v2"
)

type FileController interface {
	Upload(c *fiber.Ctx) error
}

type fileController struct {
}

func (it fileController) Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return response.Error(err, response.CodeBadRequest)
	}
	images := make([]file2.Image, 0)

	files := form.File["file"]
	for _, file := range files {
		image, err := file2.Upload(c, file)
		if err != nil {
			return response.Error(err, response.CodeBadRequest)
		}
		images = append(images, image)
	}
	return response.Success(c, images)
}

func NewFileController() FileController {
	return fileController{}
}
