package model

import "midterm-api/internal/constant"

type RequestItem struct {
	Title    string `binding:"required"`
	Amount   float64 `binding:"required,gt=0"`
	Quantity uint `binding:"required,gt=0"`
}

type RequestFindItem struct {
	//filter one status Day2
	Statuses constant.ItemStatus `form:"status"`
	//filter many statuses
	// Statuses []constant.ItemStatus `form:"status[]"`
}

type RequestUpdateItem struct {
    Status constant.ItemStatus
}

type RequestLogin struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}