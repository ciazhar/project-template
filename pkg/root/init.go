package root

import (
	"github.com/ciazhar/emobi-service/cmd/emobi-service/app"
	"github.com/ciazhar/emobi-service/pkg/root/controller/rest"
)

func Init(app app.Application) {

	controller := rest.NewRootRestController()

	r := app.Router.Group("/")
	r.Get("/", controller.Root)

}
