package usecase

import (
	"context"
	"errors"
	"github.com/ciazhar/emobi-service/internal/db"
	"github.com/ciazhar/emobi-service/pkg/vehicle/model"
	db2 "github.com/ciazhar/emobi-service/third_party/db"
	"github.com/ciazhar/emobi-service/third_party/response"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"os"
)

type VehicleUseCase interface {
	FetchVehicleCategory() ([]db.FetchVehicleCategoryRow, error)
	FetchVehicleType() ([]db.FetchVehicleTypeRow, error)
	StoreVehicle(form model.StoreVehicleForm) error
	FetchVehicle(params model.FetchParam) ([]db.FetchVehicleRow, error)
	DetailVehicle(id string) (db.DetailVehicleRow, error)
	UpsertVehiclePrice(params db.UpsertVehiclePriceParams) error
	StoreVehicleImage(params db.StoreVehicleImageParams) error
	DeleteVehiclePrice(id string) error
	UpdateVehicle(form db.UpdateVehicleParams) error
}

type vehicleUseCase struct {
	queries *db.Queries
}

func (c vehicleUseCase) UpdateVehicle(params db.UpdateVehicleParams) error {
	res, err := c.queries.ValidateVehicle(context.Background(), db.ValidateVehicleParams{
		TypeID:     params.TypeID,
		CategoryID: params.CategoryID,
	})
	if err != nil {
		return response.Error(err)
	}
	if res != "Validated" {
		return response.Error(errors.New(res))
	}

	err = c.queries.UpdateVehicle(context.Background(), params)
	return response.Error(err)
}

func (c vehicleUseCase) DeleteVehiclePrice(id string) error {
	err := c.queries.DeleteVehiclePrice(context.Background(), uuid.MustParse(id))
	return response.Error(err)
}

func (c vehicleUseCase) StoreVehicleImage(params db.StoreVehicleImageParams) error {
	//validate
	res, err := c.queries.ValidateVehicle(context.Background(), db.ValidateVehicleParams{
		VehicleID: params.VehicleID,
	})
	if err != nil {
		return response.Error(err)
	}
	if res != "Validated" {
		return response.Error(errors.New(res))
	}

	//store image
	ImagePath := "image/vehicle/"
	oldFile := viper.GetString("path") + params.Dir + params.File
	newFile := viper.GetString("path") + ImagePath + params.VehicleID.String() + "/" + params.File
	newDir := viper.GetString("path") + ImagePath + params.VehicleID.String()
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			return response.Error(err)
		}
	}
	if err := os.Rename(oldFile, newFile); err != nil {
		return response.Error(errors.New("no such file or directory"))
	}
	dirWithoutPath := ImagePath + params.VehicleID.String()
	if err := c.queries.StoreVehicleImage(context.Background(), db.StoreVehicleImageParams{
		File:      params.File,
		Dir:       dirWithoutPath,
		VehicleID: params.VehicleID,
	}); err != nil {
		return response.Error(err)
	}
	return response.Error(err)
}

func (c vehicleUseCase) UpsertVehiclePrice(params db.UpsertVehiclePriceParams) error {
	res, err := c.queries.ValidateVehicle(context.Background(), db.ValidateVehicleParams{
		VehicleID:      params.VehicleID,
		DurationTypeID: params.DurationTypeID,
	})
	if err != nil {
		return response.Error(err)
	}
	if res != "Validated" {
		return response.Error(errors.New(res))
	}
	err = c.queries.UpsertVehiclePrice(context.Background(), params)
	return response.Error(err)
}

func (c vehicleUseCase) DetailVehicle(id string) (db.DetailVehicleRow, error) {
	res, err := c.queries.DetailVehicle(context.Background(), uuid.MustParse(id))
	return res, response.Error(err)
}

func (c vehicleUseCase) FetchVehicle(params model.FetchParam) ([]db.FetchVehicleRow, error) {
	offset, limit := db2.ToOffsetLimit(params.Page, params.Size)
	res, err := c.queries.FetchVehicle(context.Background(), db.FetchVehicleParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	return res, response.Error(err)
}

func (c vehicleUseCase) FetchVehicleType() ([]db.FetchVehicleTypeRow, error) {
	res, err := c.queries.FetchVehicleType(context.Background())
	return res, response.Error(err)
}

func (c vehicleUseCase) StoreVehicle(form model.StoreVehicleForm) error {
	res, err := c.queries.ValidateVehicle(context.Background(), db.ValidateVehicleParams{
		TypeID:     form.TypeID,
		CategoryID: form.CategoryID,
	})
	if err != nil {
		return response.Error(err)
	}
	if res != "Validated" {
		return response.Error(errors.New(res))
	}

	if err := c.queries.StoreVehicle(context.Background(), db.StoreVehicleParams{
		TypeID:     form.TypeID,
		Model:      form.Model,
		Brand:      form.Brand,
		CategoryID: form.CategoryID,
		Range:      int32(form.Range),
		Wheels:     form.Wheels,
		MaxLoad:    int32(form.MaxLoad),
		TopSpeed:   int32(form.TopSpeed),
		Waterproof: int32(form.Waterproof),
		Weight:     int32(form.Weight),
	}); err != nil {
		return response.Error(err)
	}
	return nil
}

func (c vehicleUseCase) FetchVehicleCategory() ([]db.FetchVehicleCategoryRow, error) {
	res, err := c.queries.FetchVehicleCategory(context.Background())
	return res, response.Error(err)
}

func NewVehicleUseCase(queries *db.Queries) VehicleUseCase {
	return vehicleUseCase{
		queries: queries,
	}
}
