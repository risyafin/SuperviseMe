package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID      int    `json:"id"`
	CardID  int    `json:"cardId"`
	UserID  int    `json:"userId"`
	Message string `json:"message"`
}
type CommentRespons struct {
	ID      int    `json:"id"`
	CardID  int    `json:"cardId"`
	UserID  int    `json:"userId"`
	Message string `json:"message"`
}
