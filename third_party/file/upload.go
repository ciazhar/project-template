package file

import (
	"errors"
	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
)

type Image struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
}

func Upload(c *fiber.Ctx, file *multipart.FileHeader) (Image, error) {

	//response
	r := Image{}

	//content type validation
	mimeType := file.Header.Get("Content-Type")
	var contentType string
	switch mimeType {
	case "image/jpeg":
		contentType = ".jpg"
	case "image/png":
		contentType = ".png"
	default:
		return r, response.Error(errors.New("the format file is not valid"))
	}

	//create dir if not exist
	path := viper.GetString("path")
	imageDir := "image"
	newDir := path + "/" + imageDir + "/temp"
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			return r, response.Error(err)
		}
	}

	//create file name based on uuid
	id := uuid.New()
	newFileName := newDir + "/" + id.String() + contentType

	//upload file
	err := c.SaveFile(file, newFileName)
	if err != nil {
		return r, response.Error(err)
	}

	dirWithoutPath := imageDir + "/temp/"
	r.Dir = dirWithoutPath
	r.File = id.String() + contentType

	return r, nil
}
