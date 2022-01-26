package count

import "github.com/google/uuid"

type Count struct {
	Id    uuid.UUID `json:"id"`
	Label string    `json:"label"`
	Count int       `json:"count"`
}

type IntCount struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
	Count int    `json:"count"`
}
