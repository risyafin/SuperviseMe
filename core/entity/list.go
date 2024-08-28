package entity

import "gorm.io/gorm"

type List struct {
	gorm.Model
	ID       int    `json:"id"`
	ListName string `json:"listName"`
	GoalID   int    `json:"goalId"`
	Card     []Card `json:"card"`
}

type ListRespon struct {
	ID       int           `json:"id"`
	ListName string        `json:"listName"`
	GoalId   int           `json:"goalId"`
	Card     []CardRespons `json:"card"`
}
