package model

import "github.com/google/uuid"

type Machine struct {
	Id  uuid.UUID `json:"id" binding:"required"`
	Age string    `json:"age" binding:"required"`
}
