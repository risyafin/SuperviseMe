package listrepository

import (
	"superviseMe/core/entity"
	listRepo "superviseMe/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewListRepository(db *gorm.DB) listRepo.ListRepository {
	return &repository{
		DB: db,
	}
}

func (r repository) GetList(goalId string) (*entity.List, error) {
	var (
		list *entity.List
		db   = r.DB
	)

	db = db.Preload("Card")
	err := db.Where("goal_id = ?", goalId).First(&list).Error
	return list, err
}
