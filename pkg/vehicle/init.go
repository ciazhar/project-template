package vehicle

import (
	"github.com/ciazhar/emobi-service/cmd/emobi-service/app"
	"github.com/ciazhar/emobi-service/pkg/vehicle/controller/rest"
	"github.com/ciazhar/emobi-service/pkg/vehicle/usecase"
)

func Init(app app.Application) {
	uc := usecase.NewVehicleUseCase(app.DB)
	controller := rest.NewVehicleController(uc)

	r := app.Router.Group("/")
	r.Get("/vehicle-category", controller.FetchVehicleCategory)
	r.Get("/vehicle-type", controller.FetchVehicleType)
	r.Post("/vehicle", controller.StoreVehicle)
	r.Get("/vehicle", controller.FetchVehicle)
	r.Get("/vehicle/:id", controller.DetailVehicle)
	r.Post("/vehicle/price", controller.UpsertVehiclePrice)
	r.Post("/vehicle/image", controller.StoreVehicleImage)
	r.Delete("/vehicle/price/:id", controller.DeleteVehiclePrice)
	r.Put("/vehicle", controller.UpdateVehicle)
}
