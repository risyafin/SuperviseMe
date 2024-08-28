package cardrepository

import (
	"superviseMe/core/entity"
	cardRepo "superviseMe/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewCardRepository(db *gorm.DB) cardRepo.CardRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateCard(card *entity.Card) (*entity.Card, error) {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(card).Error; err != nil {
			return err
		}

		for _, checklist := range card.CheckListCard {
			checklist.CardID = card.ID
			if err := tx.Create(&checklist).Error; err != nil {
				return err
			}
		}

		return nil
	})
	return card, err
}
