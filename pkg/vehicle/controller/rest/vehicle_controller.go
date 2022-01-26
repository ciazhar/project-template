package rest

import (
	"github.com/ciazhar/emobi-service/internal/db"
	"github.com/ciazhar/emobi-service/pkg/vehicle/model"
	"github.com/ciazhar/emobi-service/pkg/vehicle/usecase"
	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/gofiber/fiber/v2"
)

type VehicleController interface {
	FetchVehicleCategory(c *fiber.Ctx) error
	FetchVehicleType(c *fiber.Ctx) error
	StoreVehicle(c *fiber.Ctx) error
	FetchVehicle(c *fiber.Ctx) error
	DetailVehicle(c *fiber.Ctx) error
	UpsertVehiclePrice(c *fiber.Ctx) error
	StoreVehicleImage(c *fiber.Ctx) error
	DeleteVehiclePrice(c *fiber.Ctx) error
	UpdateVehicle(c *fiber.Ctx) error
}

type vehicleController struct {
	VehicleUseCase usecase.VehicleUseCase
}

func (it vehicleController) UpdateVehicle(c *fiber.Ctx) error {
	var payload db.UpdateVehicleParams
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	err := it.VehicleUseCase.UpdateVehicle(payload)
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

func (it vehicleController) DeleteVehiclePrice(c *fiber.Ctx) error {
	id := c.Params("id")

	err := it.VehicleUseCase.DeleteVehiclePrice(id)
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

func (it vehicleController) StoreVehicleImage(c *fiber.Ctx) error {
	var payload db.StoreVehicleImageParams
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.VehicleUseCase.StoreVehicleImage(payload); err != nil {
		return err
	}

	return response.Success(c, nil)
}

func (it vehicleController) UpsertVehiclePrice(c *fiber.Ctx) error {
	var payload db.UpsertVehiclePriceParams
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.VehicleUseCase.UpsertVehiclePrice(payload); err != nil {
		return err
	}

	return response.Success(c, nil)
}

func (it vehicleController) DetailVehicle(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := it.VehicleUseCase.DetailVehicle(id)
	if err != nil {
		return err
	}

	return response.Success(c, res)
}

func (it vehicleController) FetchVehicle(c *fiber.Ctx) error {
	param := model.FetchParam{}
	if err := c.QueryParser(&param); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	payload, err := it.VehicleUseCase.FetchVehicle(param)
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it vehicleController) FetchVehicleType(c *fiber.Ctx) error {
	payload, err := it.VehicleUseCase.FetchVehicleType()
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it vehicleController) StoreVehicle(c *fiber.Ctx) error {
	var payload model.StoreVehicleForm
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.VehicleUseCase.StoreVehicle(payload); err != nil {
		return err
	}

	return response.Success(c, nil)
}

func (it vehicleController) FetchVehicleCategory(c *fiber.Ctx) error {
	payload, err := it.VehicleUseCase.FetchVehicleCategory()
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func NewVehicleController(VehicleUseCase usecase.VehicleUseCase) VehicleController {
	return vehicleController{VehicleUseCase: VehicleUseCase}
}
