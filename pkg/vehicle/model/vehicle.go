package model

import (
	file2 "github.com/ciazhar/emobi-service/third_party/file"
	"github.com/google/uuid"
)

type StoreVehicleForm struct {
	TypeID       uuid.UUID     `json:"type_id"`
	Model        string        `json:"model"`
	Brand        string        `json:"brand"`
	CategoryID   uuid.UUID     `json:"category_id"`
	Range        int           `json:"range"`
	Wheels       float64       `json:"wheels"`
	MaxLoad      int           `json:"max_load"`
	TopSpeed     int           `json:"top_speed"`
	Waterproof   int           `json:"waterproof"`
	Weight       int           `json:"weight"`
	VideoUrl     string        `json:"video_url"`
	PlatNumber   string        `json:"plat_number"`
	FrameNumber  string        `json:"frame_number"`
	EngineNumber string        `json:"engibe_number"`
	Image        []file2.Image `json:"image"`
}

type FetchParam struct {
	Page int `query:"page"`
	Size int `query:"size"`
}
