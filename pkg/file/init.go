package file

import (
	"github.com/ciazhar/emobi-service/cmd/emobi-service/app"
	"github.com/ciazhar/emobi-service/pkg/file/controller/rest"
)

func Init(app app.Application) {
	controller := rest.NewFileController()

	r := app.Router.Group("/")
	r.Post("/upload", controller.Upload)
}
