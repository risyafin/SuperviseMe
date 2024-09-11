package repository

import "superviseMe/core/entity"

type CardRepository interface {
	CreateCard(card *entity.Card) (*entity.Card, error) 
}
