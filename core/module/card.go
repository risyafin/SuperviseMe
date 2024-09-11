package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
	"time"
)

type CardUsecase interface {
	CreateCard(card *entity.Card) (*entity.Card, error)
}

type cardUsecase struct {
	cardRepository repository.CardRepository
}

func NewCardUsecase(cardRepository repository.CardRepository) CardUsecase {
	return &cardUsecase{cardRepository: cardRepository}
}

func (u *cardUsecase) CreateCard(card *entity.Card) (*entity.Card, error) {

	card.ListID = 1
	layout := "2006-01-02"
	card.StartDate, _ = time.Parse(layout, card.StartDate.Format(layout))
	card.EndDate, _ = time.Parse(layout, card.EndDate.Format(layout))
	card.Status = "progres"
	card.NilaiProgres = 0
	return u.cardRepository.CreateCard(card)
}
