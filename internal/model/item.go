package model

import "midterm-api/internal/constant"

type Item struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Amount   float64 `json:"amount"`
	Quantity uint `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}
type BaseResponse[DataType any] struct {
	Message string   `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}