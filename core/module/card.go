package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
)

type CardUsecase interface{
	CreateCard(card *entity.Card)(*entity.Card, error)
}

type cardUsecase struct {
	cardRepository repository.CardRepository
}

func NewCardUsecase(cardRepository repository.CardRepository) CardUsecase{
	return &cardUsecase{cardRepository: cardRepository}
}

func (u *cardUsecase) CreateCard(card *entity.Card) (*entity.Card, error){
	card.CheckListCard = append(card.CheckListCard, entity.CheckListCard{
		IsDone: "0",
		CardID: card.ID,
	})

	// 2023-12-17 00:00:00.000 format start date


	return u.cardRepository.CreateCard(card)
}