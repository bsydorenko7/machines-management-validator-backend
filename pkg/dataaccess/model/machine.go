package model

type Machine struct {
	Id  string `json:"machine" binding:"required"`
	Age string `json:"age" binding:"required"`
}
