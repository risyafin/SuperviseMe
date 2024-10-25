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
	// Mulai transaksi
	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Simpan card
	err := tx.Create(&card).Error
	if err != nil {
		tx.Rollback() // rollback jika terjadi error
		return nil, err
	}

	// Simpan checklist jika ada
	for _, checklist := range card.CheckListCard {
		checklist.CardID = card.ID // Set ID card pada checklist
		checkListCard := entity.CheckListCard{
			CardID: card.ID, // Set CardID sesuai dengan ID card yang baru saja dibuat
			Name:   "Checklist 1",
			IsDone: "0",
		}
		err := tx.Create(&checkListCard).Error
		if err != nil {
			tx.Rollback() // rollback jika terjadi error
			return nil, err
		}
	}

	// Commit transaksi jika semuanya berhasil
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return card, nil
}

func (r *repository) UpdateCheckList(name string) error {
	return r.DB.Model(&entity.CheckListCard{}).Where("name", name).Update("is_done", "1").Error
}

